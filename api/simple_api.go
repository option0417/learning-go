package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartApiServer() {
	api := gin.Default()

	api.GET("/ping", func(_context *gin.Context) {
		_context.JSON(200, gin.H{
			"messagage": "pong",
		})
	}}
	api.Run()
}
