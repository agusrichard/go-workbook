package usecases

import (
	"errors"
	"net/http"
	"twit/models"
	"twit/models/responses"
	"twit/repositories"
	"twit/utils"
)

type userUsecase struct {
	userRepository repositories.UserRepository
}

type UserUsecase interface {
	RegisterUser(user models.User) *models.RequestError
	LoginUser(userRequest models.User) (responses.LoginData, *models.RequestError)
	UserProfile(email string) (models.User, *models.RequestError)
}

func InitUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) RegisterUser(user models.User) *models.RequestError {
	if user.Email == "" || user.Password == "" {
		err := &models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Please provide email or password"),
		}
		utils.Logging(err)
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Logging(err)
		return err
	}

	user.Password = hashedPassword
	err = userUsecase.userRepository.RegisterUser(user)

	return err
}

func (userUsecase *userUsecase) LoginUser(userRequest models.User) (responses.LoginData, *models.RequestError) {
	var result responses.LoginData

	if userRequest.Email == "" || userRequest.Password == "" {
		err := &models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Please provide email and password"),
		}
		return result, err
	}
	user, err := userUsecase.userRepository.GetUserData(userRequest.Email)
	if err != nil {
		return result, err
	}

	if verified := utils.CheckPasswordHash(userRequest.Password, user.Password); !verified {
		err := &models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Wrong email or password"),
		}
		utils.Logging(err)
		return result, err
	}

	tokenStr, err := utils.GenerateToken(user)
	if err != nil {
		return result, err
	}

	loginData := responses.LoginData{
		AccessToken: tokenStr,
		User:        user,
	}
	return loginData, nil
}

func (userUsecase *userUsecase) UserProfile(email string) (models.User, *models.RequestError) {
	var result models.User
	user, err := userUsecase.userRepository.GetUserData(email)
	if err != nil {
		return result, err
	}

	return user, nil
}
