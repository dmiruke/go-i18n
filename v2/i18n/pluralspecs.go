package i18n

import "strings"

// PluralSpec defines the CLDR plural rules for a language.
// http://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html
// http://unicode.org/reports/tr35/tr35-numbers.html#Operands
type PluralSpec struct {
	Plurals    map[Plural]struct{}
	PluralFunc func(*Operands) Plural
}

// Plural returns the plural category for number as defined by
// the language's CLDR plural rules.
func (ps *PluralSpec) Plural(number interface{}) (Plural, error) {
	ops, err := newOperands(number)
	if err != nil {
		return Invalid, err
	}
	return ps.PluralFunc(ops), nil
}

func normalizePluralSpecID(id string) string {
	id = strings.Replace(id, "_", "-", -1)
	id = strings.ToLower(id)
	return id
}

func addPluralSpecs(specs map[string]*PluralSpec, ids []string, ps *PluralSpec) {
	for _, id := range ids {
		id = normalizePluralSpecID(id)
		specs[id] = ps
	}
}

func newPluralSet(plurals ...Plural) map[Plural]struct{} {
	set := make(map[Plural]struct{}, len(plurals))
	for _, plural := range plurals {
		set[plural] = struct{}{}
	}
	return set
}

func intInRange(i, from, to int64) bool {
	return from <= i && i <= to
}

func intEqualsAny(i int64, any ...int64) bool {
	for _, a := range any {
		if i == a {
			return true
		}
	}
	return false
}
