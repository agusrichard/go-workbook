package main

import (
	"twit/configs"
	"twit/handlers"
	"twit/middlewares"
	"twit/repositories"
	"twit/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configurations, database settings and auto migrations
	configModel := configs.GetConfig()
	db := configs.InitializeDB(configModel)
	configs.AutoMigrate(db)

	// Initialize repositories
	userRepository := repositories.InitUserRepository(db)

	// Initialize usecases
	userUsecase := usecases.InitUserUsecase(userRepository)

	// Initialize handlers
	userHandler := handlers.InitUserHandler(userUsecase)

	// Router for user
	router.POST("/user/register", userHandler.RegisterUser)
	router.POST("/user/login", userHandler.LoginUser)

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthenticateUser())
	{
		authorized.GET("/user/profile", userHandler.UserProfile)
	}

	router.Run(":9090")
}
