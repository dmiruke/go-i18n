package i18n

import (
	"bytes"
	"encoding"
	"strings"
	gotemplate "text/template"
)

// TODO: rename tmpl, src
type Template struct {
	tmpl *gotemplate.Template
	src  string
}

// NewTemplate returns a new template from src.
func NewTemplate(src string) (*Template, error) {
	var tmpl Template
	err := tmpl.parseTemplate(src)
	return &tmpl, err
}

func (t *Template) String() string {
	return t.src
}

// Execute executes the translation template for the given data.
func (t *Template) Execute(data interface{}) string {
	if t.tmpl == nil {
		return t.src
	}
	var buf bytes.Buffer
	if err := t.tmpl.Execute(&buf, data); err != nil {
		return err.Error()
	}
	return buf.String()
}

// MarshalText implements the TextMarshaler interface.
func (t *Template) MarshalText() ([]byte, error) {
	return []byte(t.src), nil
}

// UnmarshalText implements the TextUnmarshaler interface.
func (t *Template) UnmarshalText(src []byte) error {
	return t.parseTemplate(string(src))
}

func (t *Template) parseTemplate(src string) (err error) {
	t.src = src
	if strings.Contains(src, "{{") {
		t.tmpl, err = gotemplate.New(src).Parse(src)
	}
	return
}

var _ = encoding.TextMarshaler(&Template{})
var _ = encoding.TextUnmarshaler(&Template{})
