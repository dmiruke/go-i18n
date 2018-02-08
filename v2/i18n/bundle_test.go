package i18n_test

import (
	"reflect"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestJSON(t *testing.T) {
	var bundle i18n.Bundle
	bundle.MustParseTranslationFileBytes([]byte(`{
	"simple": "simple translation",
	"detail": {
		"description": "detail description",
		"other": "detail translation"
	},
	"everything": {
		"description": "everything description",
		"zero": "zero translation",
		"one": "one translation",
		"two": "two translation",
		"few": "few translation",
		"many": "many translation",
		"other": "other translation"
	}
}`), "en-US.json")

	expectTranslation(t, bundle, "en-US", "simple", &i18n.Translation{
		ID:    "simple",
		Other: "simple translation",
	})
	expectTranslation(t, bundle, "en-US", "detail", &i18n.Translation{
		ID:          "detail",
		Other:       "detail translation",
		Description: "detail description",
	})
	expectTranslation(t, bundle, "en-US", "everything", &i18n.Translation{
		ID:          "everything",
		Description: "everything description",
		Zero:        "zero translation",
		One:         "one translation",
		Two:         "two translation",
		Few:         "few translation",
		Many:        "many translation",
		Other:       "other translation",
	})
}

func TestYAML(t *testing.T) {
	var bundle i18n.Bundle
	bundle.MustParseTranslationFileBytes([]byte(`
# Comment
simple: simple translation

# Comment
detail:
  description: detail description 
  other: detail translation

# Comment
everything:
  description: everything description
  zero: zero translation
  one: one translation
  two: two translation
  few: few translation
  many: many translation
  other: other translation
`), "en-US.yaml")

	expectTranslation(t, bundle, "en-US", "simple", &i18n.Translation{
		ID:    "simple",
		Other: "simple translation",
	})
	expectTranslation(t, bundle, "en-US", "detail", &i18n.Translation{
		ID:          "detail",
		Other:       "detail translation",
		Description: "detail description",
	})
	expectTranslation(t, bundle, "en-US", "everything", &i18n.Translation{
		ID:          "everything",
		Description: "everything description",
		Zero:        "zero translation",
		One:         "one translation",
		Two:         "two translation",
		Few:         "few translation",
		Many:        "many translation",
		Other:       "other translation",
	})
}

func TestTOML(t *testing.T) {
	var bundle i18n.Bundle
	bundle.MustParseTranslationFileBytes([]byte(`
# Comment
simple = "simple translation"

# Comment
[detail]
description = "detail description"
other = "detail translation"

# Comment
[everything]
description = "everything description"
zero = "zero translation"
one = "one translation"
two = "two translation"
few = "few translation"
many = "many translation"
other = "other translation"
`), "en-US.toml")

	expectTranslation(t, bundle, "en-US", "simple", &i18n.Translation{
		ID:    "simple",
		Other: "simple translation",
	})
	expectTranslation(t, bundle, "en-US", "detail", &i18n.Translation{
		ID:          "detail",
		Other:       "detail translation",
		Description: "detail description",
	})
	expectTranslation(t, bundle, "en-US", "everything", &i18n.Translation{
		ID:          "everything",
		Description: "everything description",
		Zero:        "zero translation",
		One:         "one translation",
		Two:         "two translation",
		Few:         "few translation",
		Many:        "many translation",
		Other:       "other translation",
	})
}

func expectTranslation(t *testing.T, bundle i18n.Bundle, langTag, translationID string, expected *i18n.Translation) {
	actual := bundle.Translations[langTag][translationID]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("bundle.Translations[%q][%q] = %#v; want %#v", langTag, translationID, actual, expected)
	}
}
