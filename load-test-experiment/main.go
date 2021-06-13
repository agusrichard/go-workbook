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
	mdOneRp := repository.NewMediumV1Repository(db)

	// Register repositories version one
	liOneUc := usecase.NewLightV1Usecase(liOneRp)
	mdOneUc := usecase.NewMediumV1Usecase(mdOneRp)

	// Register handlers version one
	liOneHn := handler.NewLightV1Handler(liOneUc)
	mdOneHn := handler.NewMediumV1Handler(mdOneUc)


	// Register repositories version two
	liTwoRp := repository.NewLightV2Repository(db)
	mdTwoRp := repository.NewMediumV2Repository(db)

	// Register repositories version two
	liTwoUc := usecase.NewLightV2Usecase(liTwoRp)
	mdTwoUc := usecase.NewMediumV2Usecase(mdTwoRp)

	// Register handlers version two
	liTwoHn := handler.NewLightV2Handler(liTwoUc)
	mdTwoHn := handler.NewMediumV2Handler(mdTwoUc)


	// Setup gin server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Register routes
	v1 := router.Group("/v1")
	{
		v1.POST("/light/create", liOneHn.Create())
		v1.POST("/light/get", liOneHn.Get())
		v1.POST("/medium/create", mdOneHn.Create())
		v1.POST("/medium/get", mdOneHn.Get())
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/light/create", liTwoHn.Create())
		v2.POST("/light/get", liTwoHn.Get())
		v2.POST("/medium/create", mdTwoHn.Create())
		v2.POST("/medium/get", mdTwoHn.Get())
	}

	// Initialize server config and run the server
	router.Run(":9000")
}