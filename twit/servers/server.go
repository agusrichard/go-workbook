package servers

import (
	"fmt"
	"twit/configs"
	"twit/handlers"
	"twit/middlewares"
	"twit/repositories"
	"twit/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRepositories() Repositories {
	// Configurations, database settings and auto migrations
	configModel := configs.GetConfig()
	fmt.Println("configModel", configModel)
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

	// Router for user
	router.POST("/user/register", handlers.UserHandler.RegisterUser)
	router.POST("/user/login", handlers.UserHandler.LoginUser)

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthenticateUser())
	{
		authorized.GET("/user/profile", handlers.UserHandler.UserProfile)
	}

	return router
}
