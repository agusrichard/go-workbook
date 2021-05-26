package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi-tested-app/config"
	"restapi-tested-app/entities"
)


func rootHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, entities.Response{
			Success: true,
			Message: "Hello World!",
			Data: struct{}{},
		})
	}
}

func registerRoutes(router *gin.Engine, hndlrs *Handlers) {
	router.GET("/", rootHandler())
	router.POST("/tweet", hndlrs.TweetHandler.CreateTweet())
}

func SetupServer() {
	fmt.Println("Setting up server")

	configs := config.GetConfig()
	db := config.ConnectDB(configs)

	repos := setupRepositories(db)
	uscs := setupUsecases(repos)
	hdnlrs := setupHandlers(uscs)

	router := gin.Default()

	registerRoutes(router, hdnlrs)

	router.Run(":9090")
}
