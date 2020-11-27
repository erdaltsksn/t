package main

import (
	"fmt"

	"github.com/erdaltsksn/t"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	t.Configure(t.Config{
		Language:         language.Turkish,
		FallbackLanguage: language.English,
		TranslationFolder: struct {
			Path     string
			Relative bool
		}{
			Path:     "/translations",
			Relative: true,
		},
	})

	fmt.Println(t.Translate("msgHello"))
	fmt.Println(t.Translate("msgMorning", "Adem"))

	// Print the message in another language.
	msgInEnglish := message.NewPrinter(language.English).Sprintf("msgMorning", "Adam")
	fmt.Println(msgInEnglish)
}
