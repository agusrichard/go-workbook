package servers

import (
	"net/http"
	"twit/configs"
	"twit/handlers"
	"twit/repositories"
	"twit/usecases"
	"twit/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func mainHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, have a good day!",
	})
}

func SetupRepositories() (Repositories, *gorm.DB) {
	// Configurations, database settings and auto migrations
	configModel := configs.GetConfig()
	db := configs.InitializeDB(configModel)
	configs.AutoMigrate(db)

	repositories := Repositories{
		UserRepository: repositories.InitUserRepository(db),
	}

	return repositories, db
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

func SetupServer() (*gin.Engine, utils.TruncateTableExecutor) {
	router := gin.Default()

	repositories, db := SetupRepositories()
	usecases := SetupUsecases(repositories)
	handlers := SetupHandlers(usecases)

	executor := utils.InitTruncateTableExecutor(db)

	router.GET("/", mainHandler)

	// Router for user
	router.POST("/auth/register", handlers.UserHandler.RegisterUser)
	// router.POST("/auth/login", handlers.UserHandler.LoginUser)

	// authorized := router.Group("/user")
	// authorized.Use(middlewares.AuthenticateUser())
	// {
	// 	authorized.GET("/profile", handlers.UserHandler.UserProfile)
	// }

	return router, executor
}
