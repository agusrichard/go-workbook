package usecases

import "twit/repositories"

type userUsecase struct {
	userRepository repositories.UserRepository
}

type UserUsecase interface {
}

func InitUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}
