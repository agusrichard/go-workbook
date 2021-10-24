package handler

import (
	"fmt"
	"golang-restapi/model"
	"golang-restapi/repository"
	"golang-restapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Register handler function
func Register(c *gin.Context) {
	var user model.User
	var err error
	c.BindJSON(&user)
	fmt.Println("Register user", len(user.Email), len(user.Password))
	if len(user.Email) == 0 || len(user.Password) == 0 {
		utils.ResponseBadRequest(c, "Please provide email and password")
		return
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	user.UUID = uuid.New().String()
	ok := repository.CreateUser(&user)
	if ok {
		utils.ResponseSuccess(c, "Register success", gin.H{
			"email":    user.Email,
			"password": user.Password,
			"uuid":     user.UUID,
		})
	} else {
		utils.ResponseBadRequest(c, "Email has been used")
	}
}

// Login handler function
func Login(c *gin.Context) {
	var data model.LoginData
	var user model.User
	var err error
	var token string
	c.BindJSON(&data)
	if len(data.Email) == 0 || len(data.Password) == 0 {
		utils.ResponseBadRequest(c, "Please provide email and password")
		return
	}
	user, err = repository.GetUserByEmail(data.Email)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	if !user.Confirmed {
		utils.ResponseBadRequest(c, "Please confirm your account")
		return
	}

	if !utils.CheckPasswordHash(data.Password, user.Password) {
		utils.ResponseBadRequest(c, "Invalid username or password")
		return
	}
	token, err = utils.CreateToken(user.ID)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	utils.ResponseSuccess(c, "Login success", gin.H{
		"user": gin.H{
			"_id":      user.ID,
			"username": user.Email,
			"password": user.Password,
		},
		"token": token,
	})
	return
}

// ConfirmAccount ...
func ConfirmAccount(c *gin.Context) {
	var data model.ConfirmData
	var user model.User
	var err error
	var ok bool
	c.BindJSON(&data)
	if len(data.Email) == 0 || len(data.Password) == 0 || len(data.UUID) == 0 {
		utils.ResponseBadRequest(c, "Please provide email, password, and confirmation password")
	}
	user, err = repository.GetUserByEmail(data.Email)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	if user.UUID != data.UUID {
		utils.ResponseBadRequest(c, "Wrong confirmation code")
		return
	}
	ok, err = repository.ConfirmAccount(user.ID)
	if !ok || err != nil {
		utils.ResponseServerError(c)
		return
	}
	utils.ResponseSuccess(c, "Success to confirm account", gin.H{
		"email": user.Email,
	})
}

// RequestPassword ...
func RequestPassword(c *gin.Context) {
	var data model.RequestPasswordData
	var user model.User
	var err error
	var newUUID string
	var ok bool
	c.BindJSON(&data)
	if len(data.Email) == 0 {
		utils.ResponseBadRequest(c, "Please provide email")
		return
	}
	user, err = repository.GetUserByEmail(data.Email)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	if user == (model.User{}) {
		utils.ResponseNotFound(c, "User not found")
		return
	}
	newUUID = uuid.New().String()
	ok, err = repository.NewUUID(newUUID, user.Email)
	if !ok || err != nil {
		utils.ResponseServerError(c)
		return
	}
	utils.ResponseSuccess(c, "Success to get confirmation code", gin.H{
		"email":            user.Email,
		"confirmationCode": newUUID,
	})
}

// ChangePassword ...
func ChangePassword(c *gin.Context) {
	var data model.ForgotPasswordData
	var user model.User
	var err error
	var newPassword string
	var ok bool
	c.BindJSON(&data)
	if len(data.Email) == 0 || len(data.NewPassword) == 0 || len(data.UUID) == 0 {
		utils.ResponseBadRequest(c, "Please provide email, new password and confirmation code")
		return
	}
	user, err = repository.GetUserByEmail(data.Email)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	if user == (model.User{}) {
		utils.ResponseNotFound(c, "User not found")
		return
	}
	if user.UUID != data.UUID {
		utils.ResponseBadRequest(c, "Wrong confirmation code")
		return
	}
	newPassword, err = utils.HashPassword(data.NewPassword)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	ok, err = repository.ChangePassword(data.Email, newPassword)
	if !ok || err != nil {
		utils.ResponseServerError(c)
		return
	}
	utils.ResponseSuccess(c, "Success to change password", gin.H{
		"email": user.Email,
	})
}
