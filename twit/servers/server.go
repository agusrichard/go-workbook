package servers

import (
	"net/http"
	"twit/configs"
	"twit/handlers"
	"twit/middlewares"
	"twit/repositories"
	"twit/usecases"

	"github.com/gin-gonic/gin"
)

func mainHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, have a good day!",
	})
}

func SetupRepositories() Repositories {
	// Configurations, database settings and auto migrations
	configModel := configs.GetConfig()
	db := configs.InitializeDB(configModel)
	configs.AutoMigrate(db)

	repositories := Repositories{
		UserRepository: repositories.InitUserRepository(db),
	}

	return repositories
}

func SetupUsecases(repositories Repositories) Usecases {
	usecases := Usecases{
		UserUsecase: usecases.InitUserUsecase(repositories.UserRepository),
	}

	return usecases
}

func SetupHandlers(usecases Usecases) Handlers {
	handlers := Handlers{
		UserHandler: handlers.InitUserHandler(usecases.UserUsecase),
	}

	return handlers
}

func SetupServer() *gin.Engine {
	router := gin.Default()

	repositories := SetupRepositories()
	usecases := SetupUsecases(repositories)
	handlers := SetupHandlers(usecases)

	router.GET("/", mainHandler)

	// Router for user
	router.POST("/user/register", handlers.UserHandler.RegisterUser)
	router.POST("/user/login", handlers.UserHandler.LoginUser)

	authorized := router.Group("/user")
	authorized.Use(middlewares.AuthenticateUser())
	{
		authorized.GET("/profile", handlers.UserHandler.UserProfile)
	}

	return router
}
