package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"sewabuku/models"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@/sewabuku", db_username, db_password)
	connection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("could not connect database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Cart{})
	connection.AutoMigrate(&models.Book{})
	connection.AutoMigrate(&models.BookData{})
	connection.AutoMigrate(&models.Author{})
	connection.AutoMigrate(&models.Catagory{})
	connection.AutoMigrate(&models.Publisher{})
}
