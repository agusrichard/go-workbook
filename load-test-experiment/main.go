package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"load-test-experiment/config"
	"load-test-experiment/handler"
	"load-test-experiment/repository"
	"load-test-experiment/usecase"
)

func main() {
	fmt.Println("LOAD TEST EXPERIMENT")

	// Setup config and database
	configModel := config.GetConfig()
	db := config.ConnectDB(configModel)

	// Register repositories version one
	liOneRp := repository.NewLightV1Repository(db)

	// Register repositories version one
	liOneUc := usecase.NewLightV1Usecase(liOneRp)

	// Register handlers version one
	liOneHn := handler.NewLightV1Handler(liOneUc)


	// Register repositories version two
	liTwoRp := repository.NewLightV2Repository(db)

	// Register repositories version two
	liTwoUc := usecase.NewLightV2Usecase(liTwoRp)

	// Register handlers version two
	liTwoHn := handler.NewLightV2Handler(liTwoUc)


	// Setup gin server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Register routes
	v1 := router.Group("/v1")
	{
		v1.POST("/light/create", liOneHn.Create())
		v1.POST("/light/get", liOneHn.Get())
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/light/create", liTwoHn.Create())
		v2.POST("/light/get", liTwoHn.Get())
	}

	// Initialize server config and run the server
	router.Run(":9000")
}