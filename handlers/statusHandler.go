package handlers

import (
	"log"
	"net/http"
	"strings"
	"url-health/scheduler"
	"url-health/utils"

	"github.com/gin-gonic/gin"
)

// StatusSplit is a function that returns a
// which function should be used. If a query
// parameter for url is present, we want to
// do a StatusOneHandler, else  StatusAllHandler
func StatusSplit(c *gin.Context) {
	url := c.Query("url")
	if url != "" {
		StatusOneHandler(url, c)
		return
	}
	StatusAllHandler(c)
}

type StatusSingle struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

type StatusResponse struct {
	URLs []StatusSingle
}

// StatusAllHandler gets the status of all the urls in list
func StatusAllHandler(c *gin.Context) {
	resp := StatusResponse{URLs: []StatusSingle{}}
	list := scheduler.GetList()
	for url, status := range list {
		resp.URLs = append(resp.URLs, StatusSingle{URL: strings.TrimLeft(url.String(), "https://"), Status: status.String()})
	}
	c.JSON(200, resp)
}

type StatusOneResponse struct {
	URL    string `json:"url"`
	Status string `json:"status"`
}

func StatusOneHandler(url string, c *gin.Context) {
	// Check URL from request
	u, err := utils.CheckURL(strings.Trim(url, " "))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{Err: err.Error()})
		return
	}

	log.Printf("URL : %s", u.String())

	// Get the status of the URL
	status, err := scheduler.OneStatus(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{Err: err.Error()})
		return
	}
	c.JSON(200, StatusOneResponse{URL: url, Status: status.String()})
}
