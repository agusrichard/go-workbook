package repositories

import (
	"time"
	"twit/models"
	"twit/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	RegisterUser(user models.User) error
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) RegisterUser(user models.User) error {
	defer utils.MeasureExecutionTime(time.Now(), "Execution for creating user UserRepository")
	result := userRepository.db.Select("Email", "Username", "Password").Create(&user)
	// utils.Logging(result.Error, "UserRepository", "Error create user")

	return result.Error
}
