package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupServer(router *gin.Engine) {
	fmt.Println("Setting up server")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run(":9090")
}
