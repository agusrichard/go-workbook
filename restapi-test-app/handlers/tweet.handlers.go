package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi-tested-app/entities"
	"restapi-tested-app/usecases"
)

type tweetHandler struct {
	tweetUsecase usecases.TweetUsecase
}

type TweetHandler interface {
	CreateTweet() gin.HandlerFunc
}

func InitializeTweetHandler(usecase usecases.TweetUsecase) TweetHandler {
	return &tweetHandler{usecase}
}

func (handler *tweetHandler) CreateTweet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, entities.Response{
			Success: true,
			Message: "Hello World!",
			Data: struct{}{},
		})
	}
}