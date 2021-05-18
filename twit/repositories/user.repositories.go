package repositories

import (
	"errors"
	"net/http"
	"twit/models"
	"twit/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	RegisterUser(user models.User) *models.RequestError
	GetUserData(email string) (models.User, *models.RequestError)
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) RegisterUser(user models.User) *models.RequestError {
	result := userRepository.db.Select("Email", "Username", "Password").Create(&user)
	if result.Error != nil {
		err := &models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("This email has been registered. Choose another one"),
		}
		utils.Logging(err)
		return err
	}

	return nil
}

func (userRepository *userRepository) GetUserData(email string) (models.User, *models.RequestError) {
	var user models.User
	result := userRepository.db.First(&user, "email = ?", email)
	if result.Error != nil {
		err := &models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("No user found"),
		}
		utils.Logging(err)
		return models.User{}, err
	}

	return user, nil
}
