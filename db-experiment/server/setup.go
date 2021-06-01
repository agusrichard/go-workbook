package server

import (
	"db-experiment/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(router *gin.Engine, hndlrs *handlers) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "here",
			"success": true,
		})
	})
}

func SetupServer() *gin.Engine {
	fmt.Println("Setting up server")

	configs := config.GetConfig()
	db := config.ConnectDB(configs)

	repos := setupRepositories(db)
	uscs := setupUsecases(repos)
	hndlrs := setupHandlers(uscs)

	router := gin.Default()

	registerRoutes(router, hndlrs)

	return router
}
