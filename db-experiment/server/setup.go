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

	todoRoutes := router.Group("/v1/todos")
	{
		todoRoutes.POST("", h.todoHandler.CreateTodo())
		todoRoutes.GET("", h.todoHandler.GetAllTodos())
		todoRoutes.GET("/:id", h.todoHandler.GetTodoByID())
		todoRoutes.GET("/filter", h.todoHandler.FilterTodos())
		todoRoutes.PUT("", h.todoHandler.UpdateTodo())
		todoRoutes.DELETE("/:id", h.todoHandler.DeleteTodo())
	}

	todoRoutesV2 := router.Group("/v2/todos")
	{
		todoRoutesV2.POST("", h.todoHandlerV2.CreateTodo())
		todoRoutesV2.GET("", h.todoHandlerV2.GetAllTodos())
		todoRoutesV2.GET("/:id", h.todoHandlerV2.GetTodoByID())
		todoRoutesV2.GET("/filter", h.todoHandlerV2.FilterTodos())
		todoRoutesV2.PUT("", h.todoHandlerV2.UpdateTodo())
		todoRoutesV2.DELETE("/:id", h.todoHandlerV2.DeleteTodo())
	}
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
