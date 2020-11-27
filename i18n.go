package t

import (
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
	"gopkg.in/yaml.v3"
)

// dictionary is a source of translations for a single language.
type dictionary struct {
	Data map[string]string
}

// Lookup implements `Lookup()` method of 'Dictionary interface'.
// Lookup returns a message compiled with catmsg.Compile for the given key.
// It returns false for ok if such a message could not be found.
func (d *dictionary) Lookup(key string) (data string, ok bool) {
	if value, ok := d.Data[key]; ok {
		return "\x02" + value, true
	}
	return "", false
}

// Translate is used to make the translation using mapping specified in the
// `Config.TranslationFolder`. This uses `Config.Language` header if exists.
func Translate(text string, params ...interface{}) string {
	dir := getModuleDir()
	cat, err := parseYAMLDict(dir)
	if err != nil {
		log.Print(err)
	}
	message.DefaultCatalog = cat

	// return formatted text tailored to `Config.Language`.
	if acceptLanguage != nil {
		return message.NewPrinter(*acceptLanguage).Sprintf(text, params...)
	}
	return message.NewPrinter(configStore.Language).Sprintf(text, params...)
}

func getModuleDir() string {
	if configStore.TranslationFolder.Relative {
		// the argument `skip` is the number of stack frames to ascend.
		// 0 identifies the caller of Caller
		// 1 identifies this file
		// 2 identifies the called file
		_, file, _, _ := runtime.Caller(2) // `skip` argument

		return path.Join(path.Dir(file), configStore.TranslationFolder.Path)
	}

	return configStore.TranslationFolder.Path
}

func parseYAMLDict(dir string) (catalog.Catalog, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	translations := map[string]catalog.Dictionary{}
	for _, f := range files {
		yamlFile, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			return nil, err
		}

		data := map[string]string{}
		if err := yaml.Unmarshal(yamlFile, &data); err != nil {
			return nil, err
		}

		lang := strings.Split(f.Name(), ".")[0]
		translations[lang] = &dictionary{Data: data}

		// Add language to available languages.
		availableLanguages = append(availableLanguages, language.Make(lang))
	}

	// Create a Catalog from the map `translations`.
	cat, err := catalog.NewFromMap(translations)
	if err != nil {
		return nil, err
	}

	return cat, nil
}
