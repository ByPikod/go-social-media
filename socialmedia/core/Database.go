package core

import (
	"fmt"
	"socialmedia/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initializes the database
func InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"postgres",
		"root",
		"merhaba123",
		"socialmedia",
		"5432",
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	DB.AutoMigrate(
		&models.User{},
		&models.Files{},
		&models.Posts{},
		&models.Comment{},
		&models.Like{},
		&models.Friends{},
	)
}
