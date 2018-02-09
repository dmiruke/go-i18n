package i18n

import (
	"strconv"
	"strings"
	"testing"
)

type pluralTest struct {
	num    interface{}
	plural Plural
}

func runTests(t *testing.T, pluralRuleID string, tests []pluralTest) {
	pluralRuleID = normalizePluralRuleID(pluralRuleID)
	pluralRules := DefaultPluralRules()
	if rule := pluralRules[pluralRuleID]; rule != nil {
		for _, test := range tests {
			if plural, err := rule.Plural(test.num); plural != test.plural {
				t.Errorf("%s: PluralCategory(%#v) returned %s, %v; expected %s", pluralRuleID, test.num, plural, err, test.plural)
			}
		}
	} else {
		t.Errorf("could not find plural rule for locale %s", pluralRuleID)
	}

}

func appendIntegerTests(tests []pluralTest, plural Plural, examples []string) []pluralTest {
	for _, ex := range expandExamples(examples) {
		i, err := strconv.ParseInt(ex, 10, 64)
		if err != nil {
			panic(err)
		}
		tests = append(tests, pluralTest{ex, plural}, pluralTest{i, plural})
	}
	return tests
}

func appendDecimalTests(tests []pluralTest, plural Plural, examples []string) []pluralTest {
	for _, ex := range expandExamples(examples) {
		tests = append(tests, pluralTest{ex, plural})
	}
	return tests
}

func expandExamples(examples []string) []string {
	var expanded []string
	for _, ex := range examples {
		if parts := strings.Split(ex, "~"); len(parts) == 2 {
			for ex := parts[0]; ; ex = increment(ex) {
				expanded = append(expanded, ex)
				if ex == parts[1] {
					break
				}
			}
		} else {
			expanded = append(expanded, ex)
		}
	}
	return expanded
}

func increment(dec string) string {
	runes := []rune(dec)
	carry := true
	for i := len(runes) - 1; carry && i >= 0; i-- {
		switch runes[i] {
		case '.':
			continue
		case '9':
			runes[i] = '0'
		default:
			runes[i]++
			carry = false
		}
	}
	if carry {
		runes = append([]rune{'1'}, runes...)
	}
	return string(runes)
}
