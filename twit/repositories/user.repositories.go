package repositories

import (
	"twit/models"
	"twit/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	RegisterUser(user models.User) error
	GetUserData(email string) (models.User, error)
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) RegisterUser(user models.User) error {
	result := userRepository.db.Select("Email", "Username", "Password").Create(&user)
	utils.Logging(result.Error, "UserRepository", "Error create user")

	return result.Error
}

func (userRepository *userRepository) GetUserData(email string) (models.User, error) {
	var user models.User
	result := userRepository.db.Take(&user)
	utils.Logging(result.Error, "UserRepository, GetUserData", "Failed to get user data from database")

	return user, result.Error
}
