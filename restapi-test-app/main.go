package main

import (
	"github.com/gin-gonic/gin"
	"restapi-tested-app/config"
	"restapi-tested-app/server"
)

func main() {
	configs := config.GetConfig()
	config.ConnectDB(configs)

	router := gin.Default()
	server.SetupServer(router)
}