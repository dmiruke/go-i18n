package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

// Bundle stores all translations and pluralization rules.
// Generally, your application should only need a single bundle
// that is initialized early in your application's lifecycle.
type Bundle struct {
	Translations map[string]map[string]*Translation
	PluralSpecs  map[string]*PluralSpec
}

// LoadTranslationFile loads the bytes from path
// and then calls ParseTranslationFileBytes.
func (b *Bundle) LoadTranslationFile(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return b.ParseTranslationFileBytes(buf, path)
}

// MustLoadTranslationFile is similar to LoadTranslationFile
// except it panics if an error happens.
func (b *Bundle) MustLoadTranslationFile(path string) {
	if err := b.LoadTranslationFile(path); err != nil {
		panic(err)
	}
}

// LanguageTagRegex Matches language tags like en-US, and zh-Hans-CN.
// Language tags are case-insensitive.
var LanguageTagRegex = regexp.MustCompile(`[a-zA-Z]{2,}([\-_][a-zA-Z]{2,})+`)

// ParseTranslationFileBytes parses the bytes in buf to add translations to the bundle.
// It is useful for parsing translation files embedded with go-bindata.
//
// The format of the file is everything after the first ".", or the whole path if there is no ".".
// Supported formats are "json", "yaml", and "toml".
//
// The language tag of path is the last match of LanguageTagRegex.
func (b *Bundle) ParseTranslationFileBytes(buf []byte, path string) error {
	translations, err := parseTranslations(buf, path)
	if err != nil {
		return err
	}
	langTags := LanguageTagRegex.FindAllString(path, -1)
	langTag := langTags[len(langTags)-1]
	return b.AddTranslations(langTag, translations...)
}

// MustParseTranslationFileBytes is similar to ParseTranslationFileBytes
// except it panics if an error happens.
func (b *Bundle) MustParseTranslationFileBytes(buf []byte, path string) {
	if err := b.ParseTranslationFileBytes(buf, path); err != nil {
		panic(err)
	}
}

// AddTranslations adds translations for a language.
// It is useful if your translations are in a format not supported by ParseTranslationFileBytes.
func (b *Bundle) AddTranslations(langTag string, translations ...*Translation) error {
	if b.PluralSpecs == nil {
		b.PluralSpecs = DefaultPluralSpecs()
	}
	pluralID := langTag
	for i, r := range langTag {
		if r == '-' || r == '_' {
			pluralID = langTag[:i]
			break
		}
	}
	pluralSpec := b.PluralSpecs[pluralID]
	if pluralSpec == nil {
		return fmt.Errorf("no plural spec registered for %s", pluralID)
	}
	b.PluralSpecs[langTag] = pluralSpec
	if b.Translations == nil {
		b.Translations = make(map[string]map[string]*Translation)
	}
	if b.Translations[langTag] == nil {
		b.Translations[langTag] = make(map[string]*Translation)
	}
	for _, t := range translations {
		b.Translations[langTag][t.ID] = t
	}
	return nil
}

func parseTranslations(buf []byte, path string) ([]*Translation, error) {
	if len(buf) == 0 {
		return []*Translation{}, nil
	}
	var raw map[string]interface{}
	var err error
	format := path
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			format = path[i+1:]
			break
		}
	}

	switch format {
	case "json":
		err = json.Unmarshal(buf, &raw)
	case "yaml":
		err = yaml.Unmarshal(buf, &raw)
	case "toml":
		err = toml.Unmarshal(buf, &raw)
	default:
		err = fmt.Errorf("%s has unsupported format: %s", path, format)
	}
	if err != nil {
		return nil, err
	}

	var translations []*Translation
	for id, data := range raw {
		strdata := make(map[string]string)
		switch value := data.(type) {
		case string:
			strdata["other"] = value
		case map[string]interface{}:
			for k, v := range value {
				vstr, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("expected [%s][%s][%s] to be a string but got %#v", path, id, k, v)
				}
				strdata[k] = vstr
			}
		case map[interface{}]interface{}:
			for k, v := range value {
				kstr, ok := k.(string)
				if !ok {
					return nil, fmt.Errorf("[%s][%s] has a non-string key %#v", path, id, k)
				}
				vstr, ok := v.(string)
				if !ok {
					return nil, fmt.Errorf("[%s][%s][%s] has a non-string value %#v", path, id, k, v)
				}
				strdata[kstr] = vstr
			}
		default:
			return nil, fmt.Errorf("translation key %s in %s has invalid value: %#v", id, path, value)
		}
		t, err := NewTranslation(id, strdata)
		if err != nil {
			return nil, err
		}
		translations = append(translations, t)
	}
	return translations, nil
}
