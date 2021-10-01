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
			Title:         "Rich Dad Poor Dad",
			AuthorID:      3,
			PublisherID:   2,
			CategoryID:    3,
			PublisherYear: 1997,
		},
		{
			Title:         "Kambing Jantan",
			AuthorID:      4,
			PublisherID:   4,
			CategoryID:    4,
			PublisherYear: 2002,
		},
		{
			Title:         "Rumah Seribu Malaikat",
			AuthorID:      5,
			PublisherID:   3,
			CategoryID:    4,
			PublisherYear: 2009,
		},
		{
			Title:         "The Cruel Prince",
			AuthorID:      6,
			PublisherID:   5,
			CategoryID:    2,
			PublisherYear: 2009,
		},
		{
			Title:         "The Black Box",
			AuthorID:      7,
			PublisherID:   6,
			CategoryID:    2,
			PublisherYear: 2009,
		},
		{
			Title:         "Black Clover",
			AuthorID:      8,
			PublisherID:   1,
			CategoryID:    5,
			PublisherYear: 1999,
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
	}

	// ----------------
	// Category
	// ----------------

	var Category = []models.Category{
		{Name: "none"},
		{Name: "novel"},
		{Name: "motivasi"},
		{Name: "nonfiksi"},
		{Name: "comic"},
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
		{Name: "Ani", Email: "ani@g.c", Password: passStr},
		{Name: "Budi", Email: "budi@g.c", Password: passStr},
	}

	// ----------------
	// Book User mock
	// ----------------

	var BookUser = []models.Book{
		{BookDataID: 1, UserID: 1, RentPrice: 1500, Quantity: 1},
		{BookDataID: 2, UserID: 2, RentPrice: 1200, Quantity: 2},
		{BookDataID: 4, UserID: 1, RentPrice: 2400, Quantity: 1},
	}

	db.Create(&Category)
	db.Create(&Author)
	db.Create(&Publisher)
	db.Create(&bookData)
	db.Create(&User)
	db.Create(&BookUser)
}
