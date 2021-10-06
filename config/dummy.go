package config

import (
	"sewabuku/models"
	"time"

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
	// entry - add balance
	// ----------------

	var entry = []models.Entry{
		{AccountID: 2, Amount: 25000, CreatedAt: time.Now()},
		{AccountID: 4, Amount: 50000, CreatedAt: time.Now()},
		{AccountID: 6, Amount: 75000, CreatedAt: time.Now()},
	}

	// ----------------
	// Book Data
	// ----------------

	var bookData = []models.BookData{
		{
			Title:       "Rich Dad Poor Dad",
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
			Title:       "Kambing Jantan",
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
			Title:       "Rumah Seribu Malaikat",
			UserID:      5,
			Photo:       "default.jpg",
			Quantity:    1,
			AuthorID:    5,
			PublisherID: 3,
			CategoryID:  5,
			PublishYear: 2009,
			Price:       100,
		},
		{
			Title:       "The Cruel Prince",
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
			Title:       "The Black Box",
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
			Title:       "Black Clover",
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
			Title:       "Langit Bumi",
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
			Title:       "Bandung Lautan Api",
			UserID:      4,
			Quantity:    1,
			Photo:       ".jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  2,
			PublishYear: 2010,
			Price:       100,
		}, {
			Title:       "masak",
			UserID:      1,
			Quantity:    2,
			Photo:       ".jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  3,
			PublishYear: 2010,
			Price:       100,
		}, {
			Title:       "surga yang indah",
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

	db.Create(&User)
	db.Create(&entry)
	db.Create(&Category)
	db.Create(&Author)
	db.Create(&Publisher)
	db.Create(&bookData)

}
