package handler

import (
	"fmt"
	"golang-restapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserData -- Retrieve user data
func GetUserData(c *gin.Context) {
	userID := uint64(c.MustGet("userID").(float64))
	fmt.Println("GetUserData userID", userID)
	user, _ := repository.GetUserByID(userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Nice to see you bruh!",
		"data": gin.H{
			"_id":      user.ID,
			"email":    user.Email,
			"password": user.Password,
		},
	})
}
