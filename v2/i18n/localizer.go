package i18n

type Localizer struct {
	Language string
	Bundle   *Bundle
}

func (l *Localizer) Localize(id, defaultTranslation string) {

}
