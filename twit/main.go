package main

import (
	"twit/handlers"
	"twit/repositories"
	"twit/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize repositories
	userRepository := repositories.InitUserRepository()

	// Initialize usecases
	userUsecase := usecases.InitUserUsecase(userRepository)

	// Initialize handlers
	userHandler := handlers.InitUserHandler(userUsecase)

	// Router for user
	router.POST("/user/register", userHandler.RegisterUser)

	router.Run(":9090")
}
