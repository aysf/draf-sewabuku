package config

import (
	"fmt"
	"log"
	"os"
	"sewabuku/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	err := godotenv.Load("touch.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectionString :=
		fmt.Sprintf("%s:%s@/%s?parseTime=true",
			dbUsername,
			dbPassword,
			dbName,
		)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("could not connect database")
	}

	DBMigrate(db)

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.BookData{})
	db.AutoMigrate(&models.Catagory{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transfers{})
}
