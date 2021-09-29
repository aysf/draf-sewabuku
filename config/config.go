package config

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sewabuku/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	err := godotenv.Load()
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
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.BookData{})
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Catagory{})
	db.AutoMigrate(&models.Publisher{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transfers{})
}



//-------------------------------------------------------
//	DB Config for Unit Testing
//-------------------------------------------------------

func DBConnectTest() *gorm.DB {
	re := regexp.MustCompile(`^(.*` + "draf-sewabuku" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbNameTest := os.Getenv("DB_NAME_TEST")
		connectionString :=
		fmt.Sprintf("%s:%s@/%s?parseTime=true",
			dbUsername,
			dbPassword,
			dbNameTest,
		)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("could not connect database")
	}

	return db
}