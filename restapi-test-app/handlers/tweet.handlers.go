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
	GetAllTweets() gin.HandlerFunc
	CreateTweet() gin.HandlerFunc
}

func InitializeTweetHandler(usecase usecases.TweetUsecase) TweetHandler {
	return &tweetHandler{usecase}
}

func (handler *tweetHandler) GetAllTweets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tweets, err := handler.tweetUsecase.GetAllTweets()
		if err == nil {
			ctx.JSON(http.StatusOK, entities.Response{
				Success: true,
				Message: "Hello World!",
				Data: tweets,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, entities.Response{
				Success: false,
				Message: "INTERNAL SERVER ERROR",
				Data: struct{}{},
			})
		}
	}
}

func (handler *tweetHandler) CreateTweet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tweet entities.Tweet

		if err := ctx.ShouldBindJSON(&tweet); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := handler.tweetUsecase.CreateTweet(&tweet)
		if err == nil {
			ctx.JSON(http.StatusOK, entities.Response{
				Success: true,
				Message: "Hello World!",
				Data: struct{}{},
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, entities.Response{
				Success: false,
				Message: "INTERNAL SERVER ERROR",
				Data: struct{}{},
			})
		}
	}
}