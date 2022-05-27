package handlers

import (
	"fmt"
	"strings"

	"net/http"
	urlutil "net/url"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {

	// Clean the url and make sure it has https://
	url := c.Param("url")
	url = strings.TrimLeft(url, "http://")
	url = strings.TrimLeft(url, "https://")
	url = fmt.Sprintf("https://%s", url)

	parsedURL, err := urlutil.ParseRequestURI(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("URL is in incorrect form: %s", url),
		})
	}

	_ = parsedURL
}
