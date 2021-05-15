package usecases

import (
	"errors"
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
	RegisterUser(ctx *gin.Context, user models.User) *models.RequestError
	// LoginUser(ctx *gin.Context, userRequest models.User) (string, *models.RequestError)
	// UserProfile(ctx *gin.Context, email string) (models.User, *models.RequestError)
}

func InitUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) RegisterUser(ctx *gin.Context, user models.User) *models.RequestError {
	if user.Email == "" || user.Password == "" {
		err := &models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Please provide email and password"),
		}
		utils.Logging(err)
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Logging(err)
	}

	user.Password = hashedPassword
	err = userUsecase.userRepository.RegisterUser(ctx, user)

	return err
}

// func (userUsecase *userUsecase) LoginUser(ctx *gin.Context, userRequest models.User) (string, error) {
// 	if userRequest.Email == "" || userRequest.Password == "" {
// 		err := errors.New("Please provide email and password")
// 		utils.Logging(ctx, err, http.StatusBadRequest)
// 		return "", err
// 	}
// 	user, err := userUsecase.userRepository.GetUserData(ctx, userRequest.Email)

// 	if verified := utils.CheckPasswordHash(userRequest.Password, user.Password); !verified {
// 		err = errors.New("Wrong email or password")
// 		utils.Logging(ctx, err, http.StatusBadRequest)
// 		return "", err
// 	}

// 	tokenStr, err := utils.GenerateToken(user)

// 	return tokenStr, err
// }

// func (userUsecase *userUsecase) UserProfile(ctx *gin.Context, email string) (models.User, error) {
// 	user, err := userUsecase.userRepository.GetUserData(ctx, email)

// 	return user, err
// }
