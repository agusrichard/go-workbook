package main

import (
	"context"
	"fmt"
	"redis-yo/configs"
	"redis-yo/handlers"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func main() {
	rdb := configs.InitializeRedis()
	pong, _ := rdb.Ping(ctx).Result()
	fmt.Println(pong)

	router := gin.Default()

	handler := handlers.InitExampleHandler(rdb)

	router.GET("/:input", handler.MainHandler)

	router.Run(":5000")
}
