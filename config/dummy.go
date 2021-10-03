package config

import (
	"sewabuku/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertDumyData(db *gorm.DB) {

	// ----------------
	// Book Data
	// ----------------

	var bookData = []models.BookData{
		{
			Tittle:      "Rich Dad Poor Dad",
			AuthorID:    3,
			PublisherID: 2,
			CategoryID:  4,
			PublishYear: 1997,
		},
		{
			Tittle:      "Kambing Jantan",
			AuthorID:    4,
			PublisherID: 4,
			CategoryID:  5,
			PublishYear: 2002,
		},
		{
			Tittle:      "Rumah Seribu Malaikat",
			AuthorID:    5,
			PublisherID: 3,
			CategoryID:  5,
			PublishYear: 2009,
		},
		{
			Tittle:      "The Cruel Prince",
			AuthorID:    6,
			PublisherID: 5,
			CategoryID:  6,
			PublishYear: 2009,
		},
		{
			Tittle:      "The Black Box",
			AuthorID:    7,
			PublisherID: 6,
			CategoryID:  3,
			PublishYear: 2009,
		},
		{
			Tittle:      "Black Clover",
			AuthorID:    8,
			PublisherID: 1,
			CategoryID:  6,
			PublishYear: 1999,
		},
		{
			Tittle:      "Langit Bumi",
			AuthorID:    6,
			PublisherID: 3,
			CategoryID:  3,
			PublishYear: 2009,
		},
		{
			Tittle:      "Bandung Lautan Api",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  2,
			PublishYear: 2010,
		}, {
			Tittle:      "masak",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  3,
			PublishYear: 2010,
		}, {
			Tittle:      "surga yang indah",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  7,
			PublishYear: 2010,
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
	// Book User mock
	// ----------------

	var BookUser = []models.BookUser{
		{BookDataID: 1, UserID: 1, RentPrice: 1500, Quantity: 1, FileFoto: "default.jpg"},
		{BookDataID: 2, UserID: 2, RentPrice: 1200, Quantity: 2, FileFoto: "default.jpg"},
		{BookDataID: 3, UserID: 1, RentPrice: 2400, Quantity: 1, FileFoto: "default.jpg"},
		{BookDataID: 4, UserID: 4, RentPrice: 1000, Quantity: 5, FileFoto: "default.jpg"},
		{BookDataID: 5, UserID: 5, RentPrice: 1000, Quantity: 6, FileFoto: "default.jpg"},
		{BookDataID: 6, UserID: 7, RentPrice: 500, Quantity: 7, FileFoto: "default.jpg"},
		{BookDataID: 7, UserID: 5, RentPrice: 1000, Quantity: 3, FileFoto: "default.jpg"},
		{BookDataID: 8, UserID: 4, RentPrice: 200, Quantity: 2, FileFoto: "default.jpg"},
		{BookDataID: 9, UserID: 3, RentPrice: 1000, Quantity: 3, FileFoto: "default.jpg"},
		{BookDataID: 10, UserID: 1, RentPrice: 200, Quantity: 2, FileFoto: "default.jpg"},
	}

	db.Create(&Category)
	db.Create(&Author)
	db.Create(&Publisher)
	db.Create(&bookData)
	db.Create(&User)
	db.Create(&BookUser)

}
