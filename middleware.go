package t

import (
	"net/http"

	"golang.org/x/text/language"
)

// AcceptLanguageMiddleware sets `acceptLanguage` that stores language header.
func AcceptLanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if l, err := language.Parse(r.Header.Get("Accept-Language")); err == nil && inAvailableLanguages(l) {
			acceptLanguage = &l
			if acceptLanguage.IsRoot() {
				acceptLanguage = nil
			}
		} else {
			acceptLanguage = nil
		}

		next.ServeHTTP(w, r)

		return
	})
}

func inAvailableLanguages(l language.Tag) bool {
	for _, i := range availableLanguages {
		if i == l {
			return true
		}
	}

	return false
}
