package i18n_test

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TestJSON(t *testing.T) {
	var bundle i18n.Bundle
	bundle.ParseTranslationFileBytes(`
{
	"simple": "simple translation",
	"detail": {
		"translation": "detail translation",
		"description": "detail description"
	}
}`, &i18n.Language{})
}
