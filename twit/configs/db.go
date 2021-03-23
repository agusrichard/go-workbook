package configs

import (
	"fmt"
	"twit/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(config *models.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.DbName,
		config.Database.Port,
		config.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Connected to the database")
	return db
}

func AutoMigrate(db *gorm.DB) {
	// Register model and schema
	db.AutoMigrate(&models.User{})
}
