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
	liOneRp := repository.NewV1Repository(db)

	// Register repositories version one
	liOneUc := usecase.NewLightV1Usecase(liOneRp)

	// Register handlers version One
	liOneHn := handler.NewLightV1Handler(liOneUc)
	mdOneHn := handler.NewMediumV1Handler()
	hvOneHn := handler.NewHeavyV1Handler()

	// Setup gin server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Register routes
	v1 := router.Group("/v1")
	{
		v1.POST("/light", liOneHn.Handle())
		v1.POST("/medium", mdOneHn.Handle())
		v1.POST("/heavy", hvOneHn.Handle())
	}

	// Initialize server config and run the server
	router.Run(":9000")
}