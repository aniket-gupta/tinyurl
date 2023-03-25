package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {

	if url[:4] != "http" {

		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {

	if url == os.Getenv("DOMAIN") {
		return false
	}

	common := []string{"http://", "https://", "www.", ".com"}

	newURL := url

	for _, s := range common {
		newURL = strings.Replace(newURL, s, "", 1)
	}

	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")

}
