package usecases

import (
	"errors"
	"twit/models"
	"twit/repositories"
	"twit/utils"
)

type userUsecase struct {
	userRepository repositories.UserRepository
}

type UserUsecase interface {
	RegisterUser(user models.User) error
	LoginUser(userRequest models.User) (string, error)
	UserProfile(email string) (models.User, error)
}

func InitUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) RegisterUser(user models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	utils.Logging(err, "UserUsecase", "Error when hashing the password")
	user.Password = hashedPassword

	err = userUsecase.userRepository.RegisterUser(user)
	utils.Logging(err, "UserUsecase", "Error when creating the user")

	return err
}

func (userUsecase *userUsecase) LoginUser(userRequest models.User) (string, error) {
	user, err := userUsecase.userRepository.GetUserData(userRequest.Email)
	utils.Logging(err, "UserUsecase, LoginUser", "Error to get data user")

	if verified := utils.CheckPasswordHash(userRequest.Password, user.Password); !verified {
		return "", errors.New("Wrong email or password")
	}

	tokenStr, err := utils.GenerateToken(user)

	return tokenStr, err
}

func (userUsecase *userUsecase) UserProfile(email string) (models.User, error) {
	user, err := userUsecase.userRepository.GetUserData(email)
	utils.Logging(err, "UserUsecase, UserProfile", "Error to get user data")

	return user, err
}
