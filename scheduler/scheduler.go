package scheduler

import (
	"errors"
	"log"
	"net/url"
	"sync"
	"time"
)

type schedule struct {
	time  int // A measure of time in seconds to wait
	mutex sync.Mutex
}

var s schedule

type availability struct {
	url    url.URL
	status Status
}

// StartScheduler is called from configuration step
// in main, it should be ran as a goroutine
func MakeScheduler(t int) {

	// Default wait is 15 minutes
	if t > 0 {
		s = schedule{time: t}
	} else {
		s = schedule{time: 900}
	}

	log.Println("Starting scheduler")

	UpdateTime(t)
	go func() {
		for {
			RunScheduler()
		}
	}()
}

func RunScheduler() {
	// Sleep
	time.Sleep(time.Second * time.Duration(GetTime()))

	log.Println("Running update")

	// Get List of current URLs
	// Setup channel and wait group
	list := GetURLs()
	replaceMap := make(map[url.URL]Status)
	c := make(chan availability)
	var wg sync.WaitGroup

	// Spawn Go routine and pass object
	// with URLs and their status
	for _, site := range list {
		wg.Add(1)
		go func(u url.URL) {
			defer wg.Done()
			c <- availability{url: u, status: CheckURL(u)}
		}(site)
	}

	// Aggregate urls and status
	go func() {
		for a := range c {
			replaceMap[a.url] = a.status
		}
	}()

	// Wait for map to be made
	wg.Wait()
	close(c)

	UpdateStatus(replaceMap)
}

func UpdateTime(t int) error {
	if t <= 0 {
		return errors.New("Sleep must greater than 0.")
	}

	s.mutex.Lock()
	s.time = t
	s.mutex.Unlock()

	return nil
}

func GetTime() (t int) {
	s.mutex.Lock()
	t = s.time
	s.mutex.Unlock()
	return
}
