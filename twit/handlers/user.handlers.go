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
	// LoginUser(ctx *gin.Context)
	// UserProfile(ctx *gin.Context)
}

func InitUserHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase,
	}
}

func (userHandler *userHandler) RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
			Data:    struct{}{},
		})
	}

	err := userHandler.userUsecase.RegisterUser(user)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "Success to register user",
			Data:    struct{}{},
		})
	} else {
		ctx.JSON(int(err.StatusCode), models.Response{
			Success: false,
			Message: err.Error(),
			Data:    struct{}{},
		})
	}
}

// func (userHandler *userHandler) LoginUser(ctx *gin.Context) {
// 	var user models.User

// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	tokenStr, err := userHandler.userUsecase.LoginUser(ctx, user)
// 	if err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": "Success to login",
// 			"data":    tokenStr,
// 		})
// 	}
// }

// func (userHandler *userHandler) UserProfile(ctx *gin.Context) {
// 	user, err := userHandler.userUsecase.UserProfile(ctx, ctx.GetString("Email"))
// 	if err == nil {
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"message": "Success to get user profile",
// 			"data":    user,
// 		})
// 	}
// }
