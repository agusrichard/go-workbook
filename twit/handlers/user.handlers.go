package handlers

import (
	"net/http"
	"twit/models"
	"twit/models/responses"
	"twit/usecases"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUsecase usecases.UserUsecase
}

type UserHandler interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UserProfile(ctx *gin.Context)
}

func InitUserHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase,
	}
}

func (userHandler *userHandler) RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    struct{}{},
		})
	}

	err := userHandler.userUsecase.RegisterUser(user)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.Response{
			Success: true,
			Message: "Success to register user",
			Data:    struct{}{},
		})
	} else {
		ctx.JSON(int(err.StatusCode), responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    struct{}{},
		})
	}
}

func (userHandler *userHandler) LoginUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response, err := userHandler.userUsecase.LoginUser(user)
	if err == nil {
		ctx.JSON(http.StatusOK, responses.LoginUserResponse{
			Success: true,
			Message: "Success to login",
			Data:    response,
		})
	} else {
		ctx.JSON(int(err.StatusCode), responses.LoginUserResponse{
			Success: false,
			Message: err.Error(),
			Data:    response,
		})
	}
}

func (userHandler *userHandler) UserProfile(ctx *gin.Context) {
	user, err := userHandler.userUsecase.UserProfile(ctx.GetString("Email"))
	if err == nil {
		ctx.JSON(http.StatusOK, responses.Response{
			Success: true,
			Message: "Success to get user profile",
			Data:    user,
		})
	} else {
		ctx.JSON(int(err.StatusCode), responses.Response{
			Success: false,
			Message: err.Error(),
			Data:    user,
		})
	}
}
