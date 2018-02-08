package i18n

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

// TODO
func (t *Translation) Translate(plural Plural, data interface{}) string {
	var translated string
	switch plural {
	case Zero:
		translated = t.Zero
	case One:
		translated = t.One
	case Two:
		translated = t.Two
	case Few:
		translated = t.Few
	case Many:
		translated = t.Many
	case Other:
		translated = t.Other
	}
	return translated
}

type Test struct {
	Translations map[Plural]string
}

var t = &Test{
	Translations: map[Plural]string{
		Few: "World",
	},
}
