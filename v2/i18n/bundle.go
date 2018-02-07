package i18n

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"unicode"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

// Bundle stores the translations for multiple languages.
type Bundle struct {
	// DefaultLocale string
	Translations map[string]map[string]*Translation
}

// MustLoadTranslationFile is similar to LoadTranslationFile
// except it panics if an error happens.
func (b *Bundle) MustLoadTranslationFile(filename string, language *Language) {
	if err := b.LoadTranslationFile(filename, language); err != nil {
		panic(err)
	}
}

// LoadTranslationFile loads the translations from filename into memory.
func (b *Bundle) LoadTranslationFile(filename string, language *Language) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return b.ParseTranslationFileBytes(buf, language)
}

// ParseTranslationFileBytes is similar to LoadTranslationFile except it parses the bytes in buf.
//
// It is useful for parsing translation files embedded with go-bindata.
func (b *Bundle) ParseTranslationFileBytes(buf []byte, language *Language) error {
	translations, err := parseTranslations(buf)
	if err != nil {
		return err
	}
	b.AddTranslations(language, translations...)
	return nil
}

// AddTranslations adds translations for a language.
//
// It is useful if your translations are in a format not supported by LoadTranslationFile.
func (b *Bundle) AddTranslations(lang *Language, translations ...*Translation) {
	if b.Translations == nil {
		b.Translations = make(map[string]map[string]*Translation)
	}
	if b.Translations[lang.Tag] == nil {
		b.Translations[lang.Tag] = make(map[string]*Translation)
	}
	for _, t := range translations {
		b.Translations[lang.Tag][t.ID] = t
	}
}

type Translation struct {
	ID          string
	Description string
	Zero        string
	One         string
	Two         string
	Few         string
	Many        string
	Other       string
}

func parseTranslations(buf []byte) ([]*Translation, error) {
	buf = deleteLeadingComments(buf)
	if len(buf) == 0 {
		return []*Translation{}, nil
	}
	firstRune := rune(buf[0])
	var raw map[string]interface{}
	var err error
	switch firstRune {
	case '[':
		err = toml.Unmarshal(buf, raw)
	case '{':
		err = json.Unmarshal(buf, raw)
	default:
		err = yaml.Unmarshal(buf, raw)
	}
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", raw)
	return nil, nil
}

// deleteLeadingComments deletes leading newlines and comments in buf.
func deleteLeadingComments(buf []byte) []byte {
	for {
		buf = bytes.TrimLeftFunc(buf, unicode.IsSpace)
		if buf[0] == '#' {
			buf = deleteFirstLine(buf)
		} else {
			break
		}
	}
	return buf
}

// deleteLine returns buf sliced to remove the first line.
func deleteFirstLine(buf []byte) []byte {
	index := bytes.IndexRune(buf, '\n')
	if index == -1 { // If there is only one line without newline ...
		return nil // ... delete it and return nothing.
	}
	if index == len(buf)-1 { // If there is only one line with newline ...
		return nil // ... do the same as above.
	}
	return buf[index+1:]
}
