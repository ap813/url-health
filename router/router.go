package router

import (
	"url-health/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(router *gin.Engine) *gin.Engine {

	// TODO: Add config work for creating api key middleware

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	router.POST("/add", handlers.AddHandler)
	router.DELETE("/delete", handlers.RemoveHandler)
	router.GET("/list", handlers.ListHandler)

	// Splits into -> StatusAllHandler &
	// StatusOneHandler based on query string
	router.GET("/status", handlers.StatusSplit)

	router.POST("/sleep", handlers.SetSleepHandler)

	return router
}
