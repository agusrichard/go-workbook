package repositories

import (
	"net/http"
	"twit/models"
	"twit/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	RegisterUser(ctx *gin.Context, user models.User) error
	GetUserData(ctx *gin.Context, email string) (models.User, error)
}

func InitUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepository *userRepository) RegisterUser(ctx *gin.Context, user models.User) error {
	result := userRepository.db.Select("Email", "Username", "Password").Create(&user)
	utils.LogAbort(ctx, result.Error, http.StatusInternalServerError)

	return result.Error
}

func (userRepository *userRepository) GetUserData(ctx *gin.Context, email string) (models.User, error) {
	var user models.User
	result := userRepository.db.First(&user, "email = ?", email)
	utils.LogAbort(ctx, result.Error, http.StatusInternalServerError)

	return user, result.Error
}
