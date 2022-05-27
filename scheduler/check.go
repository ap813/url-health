package scheduler

import (
	"log"
	"net/http"
	"net/url"
)

// CheckURL looks for if the site at the URL is UP or DOWN
func CheckURL(url url.URL) Status {
	resp, err := http.Get(url.String())

	if err != nil {
		log.Printf("Error with call to URL (%s): %v\n", url.String(), err.Error())
		return MakeStatus(false)
	}

	if resp.StatusCode < 200 && resp.StatusCode > 299 {
		log.Printf("URL (%s) responded with non-successful code %d\n", url.String(), resp.StatusCode)
		return MakeStatus(false)
	}

	return MakeStatus(true)
}
