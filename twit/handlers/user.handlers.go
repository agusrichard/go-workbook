package handlers

import (
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

	err := userHandler.userUsecase.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success to create user",
			"data":    nil,
		})
	}
}
