package utils

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func fixURL(u string) string {
	u = strings.TrimLeft(u, "http://")
	u = strings.TrimLeft(u, "https://")
	u = strings.TrimLeft(u, "www.")
	u = fmt.Sprintf("https://www.%s", u)
	return u
}

// CheckURL is a function to check validity of URLs
func CheckURL(u string) (url.URL, error) {
	// Clean the url and make sure it has https://www.
	u = fixURL(u)

	// Verify URL passed
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return url.URL{}, errors.New(fmt.Sprintf("URL is in incorrect form: %s", u))
	}

	return *parsedURL, nil
}
