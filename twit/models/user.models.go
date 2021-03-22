package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	email    string
	username string
	password string
}
