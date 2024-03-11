package url

import (
	"fmt"
	"log"
	"net/url"
)

func MountFullUrl(lang string, term string) string {
	fullUrl := fmt.Sprintf("%s/%s", generateBaseUrl(lang), url.PathEscape(term))

	log.Println("Link to be searched: ", fullUrl)
	return fullUrl
}
