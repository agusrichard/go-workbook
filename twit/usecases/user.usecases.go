package usecases

import (
	"twit/models"
	"twit/repositories"
	"twit/utils"
)

type userUsecase struct {
	userRepository repositories.UserRepository
}

type UserUsecase interface {
	RegisterUser(user models.User) error
}

func InitUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) RegisterUser(user models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	// utils.Logging(err, "UserUsecase", "Error when hashing the password")
	user.Password = hashedPassword

	err = userUsecase.userRepository.RegisterUser(user)
	// utils.Logging(err, "UserUsecase", "Error when creating the user")

	return err
}
