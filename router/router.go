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
	router.DELETE("/delete:url", handlers.RemoveHandler)
	router.GET("/list", handlers.ListHandler)

	router.GET("/status", handlers.StatusSplit)
	// router.GET("/status", handlers.StatusAllHandler)
	// router.GET("/status/one?", handlers.StatusOneHandler)

	return router
}
