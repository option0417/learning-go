package api

import (
	"github.com/gin-gonic/gin"
)

func StartApiServer() {
	api := gin.Default()

	api.GET("/ping", func(_context *gin.Context) {
		_context.JSON(200, gin.H{
			"messagage": "pong",
		})
	})

	api.POST("/ping", func(_context *gin.Context) {
		_context.JSON(200, gin.H{
			"resp": "ok",
			"code": 0,
		})
	})
	api.Run(":9080")
}
