package server

import (
	"db-experiment/config"
	model "db-experiment/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(router *gin.Engine, h *handlers) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Hello World!",
		})
	})

	router.POST("/todos", h.todoHandler.CreateTodo())
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
