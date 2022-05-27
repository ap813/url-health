package utils

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// CheckURL is a function to check validity of URLs
func CheckURL(u string) (url.URL, error) {
	// Clean the url and make sure it has https://
	u = strings.TrimLeft(u, "http://")
	u = strings.TrimLeft(u, "https://")
	u = fmt.Sprintf("https://%s", u)

	// Verify URL passed
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return url.URL{}, errors.New(fmt.Sprintf("URL is in incorrect form: %s", u))
	}

	return *parsedURL, nil
}
