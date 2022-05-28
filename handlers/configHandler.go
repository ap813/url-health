package handlers

import (
	"net/http"
	"url-health/scheduler"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type SleepHandlerRequest struct {
	Sleep int `json:"sleep" binding:"required"`
}

func SetSleepHandler(c *gin.Context) {
	// Read JSON body
	var json SleepHandlerRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{Err: err.Error()})
		return
	}

	if err := scheduler.UpdateTime(json.Sleep); err != nil {
		c.JSON(http.StatusBadRequest, ErrorMessage{Err: err.Error()})
		return
	}

	// Return 201 created successfully
	c.Writer.WriteHeader(http.StatusCreated)
}
