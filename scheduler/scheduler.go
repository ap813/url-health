package scheduler

import (
	"net/url"
	"sync"
)

type Status string

type URLList struct {
	list map[url.URL]Status
	lock sync.RWMutex
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
	urls.lock.RLock()
	result = urls.list
	urls.lock.RUnlock()
	return
}

func SetList(newList map[url.URL]Status) {
	urls.lock.Lock()
	urls.list = newList
	urls.lock.Unlock()
}

func AddList(url url.URL, status Status) {
	urls.lock.Lock()
	urls.list[url] = status
	urls.lock.Unlock()
}
