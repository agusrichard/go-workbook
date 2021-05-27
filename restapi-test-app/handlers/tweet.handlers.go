package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi-tested-app/entities"
	"restapi-tested-app/usecases"
	"strconv"
)

type tweetHandler struct {
	tweetUsecase usecases.TweetUsecase
}

type TweetHandler interface {
	GetAllTweets() gin.HandlerFunc
	GetTweetByID() gin.HandlerFunc
	SearchTweetByText() gin.HandlerFunc
	CreateTweet() gin.HandlerFunc
	UpdateTweet() gin.HandlerFunc
	DeleteTweet() gin.HandlerFunc
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

func (handler *tweetHandler) GetTweetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		tweet, err := handler.tweetUsecase.GetTweetByID(id)
		if err == nil {
			ctx.JSON(http.StatusOK, entities.Response{
				Success: true,
				Message: "Hello World!",
				Data: tweet,
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

func (handler *tweetHandler) SearchTweetByText() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchParam := ctx.Query("search")

		tweets, err := handler.tweetUsecase.SearchTextByText(searchParam)
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

func (handler *tweetHandler) UpdateTweet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tweet entities.Tweet

		if err := ctx.ShouldBindJSON(&tweet); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := handler.tweetUsecase.UpdateTweet(&tweet)
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

func (handler *tweetHandler) DeleteTweet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		err = handler.tweetUsecase.DeleteTweet(id)
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