package config

import (
	"sewabuku/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertDumyData(db *gorm.DB) {

	// ----------------
	// User mock
	// ----------------

	pass, _ := bcrypt.GenerateFromPassword([]byte("123"), 14)
	passStr := string(pass)
	var User = []models.User{
		{Name: "Ani", Email: "ani@g.c", Password: passStr, Address: "jakarta"},
		{Name: "Budi", Email: "budi@g.c", Password: passStr, Address: "depok"},
		{Name: "Danang", Email: "danang@.com", Password: passStr, Address: "bandung"},
		{Name: "Sanas", Email: "sanas@gmail.", Password: passStr, Address: "sumedang"},
		{Name: "kokoh", Email: "kooko@.com", Password: passStr, Address: "cikarang"},
		{Name: "Saness", Email: "saneess@gmail.", Password: passStr, Address: "surabaya"},
		{Name: "kokoh", Email: "koiooako@.com", Password: passStr, Address: "palembang"},
	}

	// ----------------
	// Book Data
	// ----------------

	var bookData = []models.BookData{
		{
			Tittle:      "Rich Dad Poor Dad",
			UserID:      2,
			Quantity:    2,
			Photo:       "default.jpg",
			AuthorID:    3,
			PublisherID: 2,
			CategoryID:  4,
			PublishYear: 1997,
			Price:       100,
		},
		{
			Tittle:      "Kambing Jantan",
			UserID:      4,
			Photo:       "default.jpg",
			Quantity:    2,
			AuthorID:    4,
			PublisherID: 4,
			CategoryID:  5,
			PublishYear: 2002,
			Price:       100,
		},
		{
			Tittle:      "Rumah Seribu Malaikat",
			UserID:      5,
			Photo:       "default.jpg",
			Quantity:    2,
			AuthorID:    5,
			PublisherID: 3,
			CategoryID:  5,
			PublishYear: 2009,
			Price:       100,
		},
		{
			Tittle:      "The Cruel Prince",
			UserID:      2,
			Photo:       ".jpg",
			Quantity:    0,
			AuthorID:    6,
			PublisherID: 5,
			CategoryID:  6,
			PublishYear: 2009,
			Price:       100,
		},
		{
			Tittle:      "The Black Box",
			UserID:      1,
			Quantity:    1,
			Photo:       ".jpg",
			AuthorID:    7,
			PublisherID: 6,
			CategoryID:  3,
			PublishYear: 2009,
			Price:       200,
		},
		{
			Tittle:      "Black Clover",
			UserID:      1,
			AuthorID:    8,
			Quantity:    1,
			Photo:       ".jpg",
			PublisherID: 1,
			CategoryID:  6,
			PublishYear: 1999,
			Price:       100,
		},
		{
			Tittle:      "Langit Bumi",
			UserID:      2,
			Quantity:    1,
			Photo:       "jpg",
			AuthorID:    6,
			PublisherID: 3,
			CategoryID:  3,
			PublishYear: 2009,
			Price:       10,
		},
		{
			Tittle:      "Bandung Lautan Api",
			UserID:      4,
			Quantity:    1,
			Photo:       ".jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  2,
			PublishYear: 2010,
			Price:       100,
		}, {
			Tittle:      "masak",
			UserID:      1,
			Quantity:    2,
			Photo:       ".jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  3,
			PublishYear: 2010,
			Price:       100,
		}, {
			Tittle:      "surga yang indah",
			UserID:      1,
			Quantity:    1,
			Photo:       ".jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  7,
			PublishYear: 2010,
			Price:       100,
		},
	}

	// ----------------
	// Author
	// ----------------

	var Author = []models.Author{
		{Name: "none"},
		{Name: "JK. Rowling"},
		{Name: "Robert Kiyosaki"},
		{Name: "Raditya Dika"},
		{Name: "Yuli Badawi"},
		{Name: "Holly Black"},
		{Name: "Michael, Connelly"},
		{Name: "Tabata & Yuuki"},
		{Name: "Tere liye"},
	}

	// ----------------
	// Category
	// ----------------

	var Category = []models.Category{
		{Name: "none"},
		{Name: "sejarah"},
		{Name: "novel"},
		{Name: "motivasi"},
		{Name: "non fiksi"},
		{Name: "comic"},
		{Name: "Agama"},
		{Name: "horror"},
	}

	// ----------------
	// Publisher
	// ----------------

	var Publisher = []models.Publisher{
		{Name: "none"},
		{Name: "gramedia"},
		{Name: "mizan"},
		{Name: "andi"},
		{Name: "LBYR"},
		{Name: "Orion"},
	}

	// ----------------
	// Book User mock
	// ----------------

	// var BookUser = []models.BookUser{
	// 	{BookDataID: 1, UserID: 1, RentPrice: 1500, Quantity: 1, FileFoto: "default.jpg"},
	// 	{BookDataID: 2, UserID: 2, RentPrice: 1200, Quantity: 2, FileFoto: "default.jpg"},
	// 	{BookDataID: 3, UserID: 1, RentPrice: 2400, Quantity: 1, FileFoto: "default.jpg"},
	// 	{BookDataID: 4, UserID: 4, RentPrice: 1000, Quantity: 5, FileFoto: "default.jpg"},
	// 	{BookDataID: 5, UserID: 5, RentPrice: 1000, Quantity: 6, FileFoto: "default.jpg"},
	// 	{BookDataID: 6, UserID: 7, RentPrice: 500, Quantity: 7, FileFoto: "default.jpg"},
	// 	{BookDataID: 7, UserID: 5, RentPrice: 1000, Quantity: 3, FileFoto: "default.jpg"},
	// 	{BookDataID: 8, UserID: 4, RentPrice: 200, Quantity: 2, FileFoto: "default.jpg"},
	// 	{BookDataID: 9, UserID: 3, RentPrice: 1000, Quantity: 3, FileFoto: "default.jpg"},
	// 	{BookDataID: 10, UserID: 1, RentPrice: 200, Quantity: 2, FileFoto: "default.jpg"},
	// }

	db.Create(&User)
	db.Create(&Category)
	db.Create(&Author)
	db.Create(&Publisher)
	db.Create(&bookData)

}
