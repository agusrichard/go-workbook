package repository

import (
	"fmt"
	"golang-restapi/config"
	"golang-restapi/model"
)

// CreateUser -- Create user
func CreateUser(user *model.User) bool {
	sqlQuery := `
		INSERT INTO users (email, password, uuid) 
		VALUES ($1, $2, $3);
	`

	_, err := config.DB.Exec(sqlQuery, user.Email, user.Password, user.UUID)
	if err != nil {
		// panic(err)
		return false
	}
	fmt.Printf("Success to create user %v\n", *user)
	return true
}

// GetUserByEmail -- Get user by Email
func GetUserByEmail(email string) (model.User, error) {
	sqlQuery := `
		SELECT _id, email, password, uuid, confirmed FROM users
		WHERE email = $1;
	`
	rows, err := config.DB.Query(sqlQuery, email)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()
	var user model.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.UUID,
			&user.Confirmed,
		)
		if err != nil {
			return model.User{}, err
		}
	}
	err = rows.Err()
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetUserByID -- Get user data
func GetUserByID(userID uint64) (model.User, error) {
	sqlQuery := `
		SELECT _id, email, password FROM users
		WHERE _id = $1
	`
	rows, err := config.DB.Query(sqlQuery, userID)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()
	var user model.User
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return model.User{}, err
		}
	}
	err = rows.Err()
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// ConfirmAccount --  Update confirm field
func ConfirmAccount(userID uint64) (bool, error) {
	sqlQuery := `
		UPDATE users
		SET confirmed=TRUE
		WHERE _id=$1
	`

	_, err := config.DB.Query(sqlQuery, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewUUID --
func NewUUID(uuid, email string) (bool, error) {
	sqlQuery := `
		UPDATE users
		SET uuid=$1
		WHERE email=$2
	`

	_, err := config.DB.Query(sqlQuery, uuid, email)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ChangePassword -- Forgot Password
func ChangePassword(email string, newPassword string) (bool, error) {
	sqlQuery := `
		UPDATE users
		SET password=$1
		WHERE email=$2
	`

	_, err := config.DB.Query(sqlQuery, newPassword, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
