package t

import (
	"golang.org/x/text/language"
)

// Config stores all configurations for `t` package.
type Config struct {
	Language          language.Tag
	FallbackLanguage  language.Tag
	TranslationFolder struct {
		Path     string
		Relative bool
	}
}

var (
	configStore        Config
	acceptLanguage     *language.Tag
	availableLanguages []language.Tag
)

// Configure configures the `t` package.
func Configure(c Config) {
	configStore = c
}
