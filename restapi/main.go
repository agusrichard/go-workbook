package main

import (
	"golang-restapi/config"
	"golang-restapi/handler"
	"golang-restapi/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	// Initialize database connection
	config.InitDb()

	// Routes for authentication
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", handler.Register)
		authRoute.POST("/login", handler.Login)
		authRoute.POST("/confirm", handler.ConfirmAccount)
		authRoute.POST("/forgot-password", handler.RequestPassword)
		authRoute.POST("/forgot-password/confirm", handler.ChangePassword)
	}

	// Routes for User
	userRoute := r.Group("/user")
	userRoute.Use(middleware.AuthMiddleware())
	{
		userRoute.GET("/", handler.GetUserData)
	}

	// Routes for Service
	serviceRoute := r.Group("/service")
	serviceRoute.Use(middleware.AuthMiddleware())
	{
		serviceRoute.POST("/", handler.CreateServiceRequest)
		serviceRoute.GET("/", handler.GetServices)
	}

	r.Run()
}
