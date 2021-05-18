package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type exampleHandler struct {
	rdb *redis.Client
}

type ExampleHandler interface {
	MainHandler(ctx *gin.Context)
}

func InitExampleHandler(rdb *redis.Client) ExampleHandler {
	return &exampleHandler{
		rdb,
	}
}

func (handler *exampleHandler) MainHandler(ctx *gin.Context) {
	input := ctx.Param("input")

	existanceStatus, _ := handler.rdb.Exists(ctx, input).Result()

	if existanceStatus == 0 {
		handler.rdb.Set(ctx, input, "Sekardayu Hana Pradiani", time.Minute).Result()
		MockingSomeOperation()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"echo":    input,
			"data":    "",
		})
	} else {
		result, _ := handler.rdb.Get(ctx, input).Result()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"echo":    input,
			"data":    result,
		})
	}
}

func MockingSomeOperation() {
	fmt.Println("Start sleeping")
	time.Sleep(time.Second * 5)
	fmt.Println("End sleeping")
}
