package servers

import (
	"net/http"
	"twit/configs"
	"twit/handlers"
	"twit/repositories"
	"twit/usecases"

	"github.com/gin-gonic/gin"
)

func mainHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, have a good day!",
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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

	router.Use(CORSMiddleware())

	router.GET("/", mainHandler)

	// Router for user
	router.POST("/auth/register", handlers.UserHandler.RegisterUser)
	// router.POST("/auth/login", handlers.UserHandler.LoginUser)

	// authorized := router.Group("/user")
	// authorized.Use(middlewares.AuthenticateUser())
	// {
	// 	authorized.GET("/profile", handlers.UserHandler.UserProfile)
	// }

	return router
}
