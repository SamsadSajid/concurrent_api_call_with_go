package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Controller() *gin.Engine{
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "message",
		})
	})


	router.GET("/service/:userName", getUserInfo)

	return router
}

