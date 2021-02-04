package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"grpc-auth/models"
	"grpc-auth/repository"
	"grpc-auth/utils"
	"log"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	Register(username, password string) (bool, error)
	Login(username, password string) (string, error)
	ValidateToken(token string) (string, error)
}

func InitUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (userUsecase *userUsecase) Register(username, password string) (bool, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("Error to register user: ", err)
		return false, err
	}

	_, err = userUsecase.userRepository.CreateUser(username, hashedPassword)
	if err != nil {
		log.Println("Error to register user: ", err)
		return false, err
	}

	return true, nil
}

func (userUsecase *userUsecase) Login(username, password string) (string, error) {
	user, err := userUsecase.userRepository.GetUserByUsername(username)
	if err != nil {
		log.Println("Error to get user by username", err)
	}
	ok := utils.CheckPasswordHash(password, user.Password)
	if !ok {
		log.Println("Wrong password")
		return "", errors.New("Wrong password")
	}
	tokenString, err := utils.GenerateToken(models.User{
		ID:       user.ID,
		Username: user.Username,
	})
	if err != nil {
		log.Println("Error to login", err)
	}
	return tokenString, nil
}

func (userUsecase *userUsecase) ValidateToken(token string) (string, error) {
	user, err := utils.ParseToken(token)
	fmt.Println("user", user)
	if err != nil {
		log.Println("Error to validate token", err)
		return "", err
	}
	result, err := json.Marshal(&user)
	if err != nil {
		log.Println("Error to validate token", err)
		return "", err
	}
	return string(result), nil
}
