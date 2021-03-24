package usecases

import (
	"errors"
	"fmt"
	"net/http"
	"twit/models"
	"twit/repositories"
	"twit/utils"

	"github.com/gin-gonic/gin"
)

type userUsecase struct {
	userRepository repositories.UserRepository
}

type UserUsecase interface {
	RegisterUser(ctx *gin.Context, user models.User) error
	LoginUser(ctx *gin.Context, userRequest models.User) (string, error)
	UserProfile(ctx *gin.Context, email string) (models.User, error)
}

func InitUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) RegisterUser(ctx *gin.Context, user models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	utils.LogAbort(ctx, err, http.StatusInternalServerError)
	user.Password = hashedPassword

	err = userUsecase.userRepository.RegisterUser(ctx, user)

	return err
}

func (userUsecase *userUsecase) LoginUser(ctx *gin.Context, userRequest models.User) (string, error) {
	fmt.Println("userRequest", userRequest.Email)
	user, err := userUsecase.userRepository.GetUserData(ctx, userRequest.Email)
	fmt.Println("user", user)

	if verified := utils.CheckPasswordHash(userRequest.Password, user.Password); !verified {
		err = errors.New("Wrong email or password")
		utils.LogAbort(ctx, err, http.StatusBadRequest)
		return "", err
	}

	tokenStr, err := utils.GenerateToken(user)

	return tokenStr, err
}

func (userUsecase *userUsecase) UserProfile(ctx *gin.Context, email string) (models.User, error) {
	user, err := userUsecase.userRepository.GetUserData(ctx, email)
	utils.LogAbort(ctx, err, http.StatusInternalServerError)

	return user, err
}
