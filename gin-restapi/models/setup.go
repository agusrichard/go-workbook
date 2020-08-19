package models

import (
	"github.com/jinzhu/gorm"
)

// DB -- Database
var DB *gorm.DB

func connectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})
}
