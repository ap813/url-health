package scheduler

import (
	"errors"
	"net/url"
	"sync"
)

type Status string

func (s Status) String() string {
	return string(s)
}

type URLList struct {
	list  map[url.URL]Status
	mutex sync.RWMutex
}

var urls URLList

func MakeStatus(b bool) Status {
	if b {
		return Status("UP")
	} else {
		return Status("DOWN")
	}
}

// GetList: get map of urls and status
func GetList() (result map[url.URL]Status) {
	result = make(map[url.URL]Status)
	urls.mutex.RLock()

	// Make a copy of the map to return
	for key, val := range urls.list {
		result[key] = val
	}

	urls.mutex.RUnlock()
	return
}

// GetURLs produces a copy of the urls being tracked
func GetURLs() (result []url.URL) {
	result = []url.URL{}
	urls.mutex.RLock()

	// Make a copy of the map to return
	for key := range urls.list {
		result = append(result, key)
	}

	urls.mutex.RUnlock()
	return
}

// SetList will replace the entire map of urls and their status
func SetList(newList map[url.URL]Status) {
	urls.mutex.Lock()
	urls.list = newList
	urls.mutex.Unlock()
}

// AddList inserts a new url and status to the map
func AddList(url url.URL, status Status) {
	urls.mutex.Lock()
	urls.list[url] = status
	urls.mutex.Unlock()
}

// DeleteURL removes a URL from the map if it exists
func DeleteURL(url url.URL) {
	urls.mutex.Lock()
	delete(urls.list, url)
	urls.mutex.Unlock()
}

// OneStatus will get the status of a valid url or send an error
func OneStatus(url url.URL) (Status, error) {
	urls.mutex.RLock()
	status, ok := urls.list[url]
	urls.mutex.RUnlock()
	if !ok {
		return "", errors.New("URL not present in list")
	}
	return status, nil
}

// UpdateStatus will replace status of existing URLs
func UpdateStatus(updateList map[url.URL]Status) {
	urls.mutex.Lock()
	for key, val := range updateList {
		// Only values still in urls.list
		if _, ok := urls.list[key]; ok {
			urls.list[key] = val
		}
	}
	urls.mutex.Unlock()
}
