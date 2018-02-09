package i18n

// Language is a written human language.
type Language struct {
	// Tag uniquely identifies the language as defined by RFC 5646.
	//
	// Most language tags are a two character language code (ISO 639-1)
	// optionally followed by a dash and a two character country code (ISO 3166-1).
	// (e.g. en, pt-br)
	Tag string
	*PluralRule
}
