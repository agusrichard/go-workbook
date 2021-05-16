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
	// GetUserData(ctx *gin.Context, email string) (models.User, error)
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

// func (userRepository *userRepository) GetUserData(ctx *gin.Context, email string) (models.User, error) {
// 	var user models.User
// 	result := userRepository.db.First(&user, "email = ?", email)
// 	utils.Logging(ctx, result.Error, http.StatusInternalServerError)

// 	return user, result.Error
// }
