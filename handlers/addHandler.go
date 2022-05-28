package handlers

import (
	"url-health/scheduler"
	"url-health/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type AddHandlerRequest struct {
	URL string `json:"url" binding:"required"`
}

func AddHandler(c *gin.Context) {
	// Read JSON body
	var json AddHandlerRequest
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

	// Check status then write status to list
	status := scheduler.CheckURL(parsedURL)
	scheduler.AddList(parsedURL, status)

	// Return 201 created successfully
	c.Writer.WriteHeader(http.StatusCreated)
}
