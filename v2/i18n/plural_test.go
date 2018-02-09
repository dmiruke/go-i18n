package i18n

import (
	"testing"
)

func TestNewPluralForm(t *testing.T) {
	tests := []struct {
		src        string
		pluralForm PluralForm
		err        bool
	}{
		{"zero", Zero, false},
		{"one", One, false},
		{"two", Two, false},
		{"few", Few, false},
		{"many", Many, false},
		{"other", Other, false},
		{"asdf", Invalid, true},
	}
	for _, test := range tests {
		pluralForm, err := NewPluralForm(test.src)
		wrongErr := (err != nil && !test.err) || (err == nil && test.err)
		if pluralForm != test.pluralForm || wrongErr {
			t.Errorf("NewPlural(%#v) returned %#v,%#v; expected %#v", test.src, pluralForm, err, test.pluralForm)
		}
	}
}
