package handlers

import (
	"net/http"
	"url-health/scheduler"
	"url-health/utils"

	"github.com/gin-gonic/gin"
)

type RemoveHandlerRequest struct {
	URL string `json:"url" binding:"required"`
}

// RemoveHandler will take a URL out of the list
func RemoveHandler(c *gin.Context) {
	// Read JSON body
	var json RemoveHandlerRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{Err: err.Error()})
		return
	}

	// Clean the url and make sure it has https://
	parsedURL, err := utils.CheckURL(json.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{Err: err.Error()})
		return
	}

	scheduler.DeleteURL(parsedURL)

	// Add to response body and return 201
	c.Writer.WriteHeader(http.StatusNoContent)
}
