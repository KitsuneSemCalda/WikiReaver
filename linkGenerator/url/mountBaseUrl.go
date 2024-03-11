package url

import "fmt"

func generateBaseUrl(lang string) string {
	if lang == "" {
		lang = "en"
	}

	return fmt.Sprintf("https://%s.wikipedia.org/wiki", lang)
}
