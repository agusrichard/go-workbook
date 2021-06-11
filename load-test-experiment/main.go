package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"load-test-experiment/config"
	"load-test-experiment/handler"
	"net/http"
	"time"
)

func main() {
	fmt.Println("LOAD TEST EXPERIMENT")

	// Setup config and database
	configModel := config.GetConfig()
	config.ConnectDB(configModel)

	// Register handlers Version One
	liOneHn := handler.NewLightV1Handler()
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
	s := &http.Server{
		Addr:           ":9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}