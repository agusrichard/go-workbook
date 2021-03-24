package handlers

import (
	"fmt"
	"net/http"
	"twit/models"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := userHandler.userUsecase.RegisterUser(ctx, user)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success to register user",
			"data":    nil,
		})
	}
}

func (userHandler *userHandler) LoginUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println("user", user)
	tokenStr, err := userHandler.userUsecase.LoginUser(ctx, user)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success to login",
			"data":    tokenStr,
		})
	}
}

func (userHandler *userHandler) UserProfile(ctx *gin.Context) {
	user, _ := userHandler.userUsecase.UserProfile(ctx, ctx.GetString("email"))
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success to get user profile",
		"data":    user,
	})
}
