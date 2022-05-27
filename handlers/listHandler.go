package handlers

import (
	"strings"
	"url-health/scheduler"

	"github.com/gin-gonic/gin"
)

type ListHandlerReturn struct {
	URLs []string `json:"urls"`
}

func ListHandler(c *gin.Context) {
	result := ListHandlerReturn{URLs: []string{}}
	list := scheduler.GetURLs()
	for _, val := range list {
		result.URLs = append(result.URLs, strings.TrimLeft(val.String(), "https://"))
	}
	c.JSON(200, result)
}
