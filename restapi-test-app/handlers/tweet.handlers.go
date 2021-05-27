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
	GetAllTweets(ctx *gin.Context) *entities.AppResult
	GetTweetByID() gin.HandlerFunc
	SearchTweetByText() gin.HandlerFunc
	CreateTweet(ctx *gin.Context) *entities.AppResult
	UpdateTweet() gin.HandlerFunc
	DeleteTweet() gin.HandlerFunc
}

func InitializeTweetHandler(usecase usecases.TweetUsecase) TweetHandler {
	return &tweetHandler{usecase}
}

func (handler *tweetHandler) GetAllTweets(*gin.Context) *entities.AppResult {
	var result entities.AppResult

	tweets, err := handler.tweetUsecase.GetAllTweets()
	if err == nil {
		result.StatusCode = http.StatusOK
		result.Message = "Success to get all tweets"
		if len(*tweets) == 0 {
			result.Data = []interface{}{}
		} else {
			result.Data = tweets
		}
	} else {
		result.StatusCode = http.StatusInternalServerError
		result.Err = err
		result.Data = []interface{}{}
	}

	return &result
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

		tweets, err := handler.tweetUsecase.SearchTweetByText(searchParam)
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

func (handler *tweetHandler) CreateTweet(ctx *gin.Context) *entities.AppResult {
	var tweet entities.Tweet
	var result entities.AppResult

	if err := ctx.ShouldBindJSON(&tweet); err != nil {
		result.Err = err
		result.Message = "username and text can not be empty"
		result.StatusCode = http.StatusBadRequest
		return &result
	}

	err := handler.tweetUsecase.CreateTweet(&tweet)
	if err == nil {
		result.Message = "Success to create tweet"
		result.StatusCode = http.StatusCreated
	} else {
		result.Err = err.Err
		result.Message = err.Err.Error()
		result.StatusCode = err.StatusCode
	}

	return &result
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