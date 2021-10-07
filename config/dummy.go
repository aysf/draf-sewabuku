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
		{Name: "Ani", Email: "ani@mail.com", Password: passStr, Address: "jakarta"},
		{Name: "Baiq", Email: "baiq@mail.com", Password: passStr, Address: "depok"},
		{Name: "Citra", Email: "citra@mail.com", Password: passStr, Address: "depok"},
		{Name: "Desi", Email: "desi@mail.com", Password: passStr, Address: "bandung"},
		{Name: "Eko", Email: "eko@mail.com", Password: passStr, Address: "bandung"},
		{Name: "Fakhry", Email: "fakhry@mail.com", Password: passStr, Address: "sumedang"},
		{Name: "Gun Gun", Email: "gun@mail.com", Password: passStr, Address: "cikarang"},
		{Name: "Heri", Email: "heri@mail.com", Password: passStr, Address: "surabaya"},
	}

	// ----------------
	// entry - add balance
	// ----------------

	var entry = []models.Entry{
		{AccountID: 5, Amount: 25000, CreatedAt: time.Now()},
		{AccountID: 6, Amount: 50000, CreatedAt: time.Now()},
		{AccountID: 7, Amount: 75000, CreatedAt: time.Now()},
		{AccountID: 8, Amount: 100000, CreatedAt: time.Now()},
	}

	// ----------------
	// Book Data
	// ----------------

	var bookData = []models.BookData{
		{
			Title:       "Rich Dad Poor Dad",
			UserID:      1,
			Quantity:    5,
			Photo:       "default.jpg",
			AuthorID:    3,
			PublisherID: 2,
			CategoryID:  4,
			PublishYear: 1997,
			Price:       100,
		},
		{
			Title:       "Kambing Jantan",
			UserID:      2,
			Photo:       "default.jpg",
			Quantity:    1,
			AuthorID:    4,
			PublisherID: 4,
			CategoryID:  5,
			PublishYear: 2002,
			Price:       100,
		},
		{
			Title:       "Rumah Seribu Malaikat",
			UserID:      1,
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
			UserID:      1,
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
			UserID:      3,
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
			UserID:      4,
			AuthorID:    8,
			Quantity:    1,
			Photo:       ".jpg",
			PublisherID: 1,
			CategoryID:  6,
			PublishYear: 1999,
			Price:       100,
		},
		{
			Title:       "La Tahzan",
			UserID:      6,
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
			Title:       "Juru Masak Para Maiko 05",
			UserID:      1,
			Quantity:    2,
			Photo:       "default.jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  3,
			PublishYear: 2010,
			Price:       100,
		}, {
			Title:       "surga yang indah",
			UserID:      1,
			Quantity:    1,
			Photo:       "default.jpg",
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

	var Cart = []models.Cart{
		{ID: 1, UserID: 1, BookDataID: 4, DateLoan: time.Now(), DateDue: time.Now(), DateReturn: time.Now(), IsReturn: true},
		{ID: 2, UserID: 3, BookDataID: 4, DateLoan: time.Now(), DateDue: time.Now(), DateReturn: time.Now(), IsReturn: true},
		{ID: 3, UserID: 2, BookDataID: 3, DateLoan: time.Now(), DateDue: time.Now(), DateReturn: time.Now(), IsReturn: true},
		{ID: 4, UserID: 2, BookDataID: 1, DateLoan: time.Now(), DateDue: time.Now(), DateReturn: time.Now(), IsReturn: true},
		{ID: 5, UserID: 1, BookDataID: 2, DateLoan: time.Now(), DateDue: time.Now(), DateReturn: time.Now(), IsReturn: true},
		{ID: 6, UserID: 3, BookDataID: 3, DateLoan: time.Now(), DateDue: time.Now(), DateReturn: time.Now(), IsReturn: true},
	}

	var Rating = []models.Rating{
		{ID: 1, CartID: 1, RateBook: 8, RateBorrower: 7, DescRateBook: "asala", DescRateBorrower: "gtddnj"},
		{ID: 2, CartID: 2, RateBook: 7, RateBorrower: 7, DescRateBook: "asali", DescRateBorrower: "gtddnj"},
		{ID: 3, CartID: 3, RateBook: 7, RateBorrower: 7, DescRateBook: "asalcoba", DescRateBorrower: "gtddnj"},
		{ID: 4, CartID: 4, RateBook: 8, RateBorrower: 7, DescRateBook: "asalaja", DescRateBorrower: "gtddnj"},
		{ID: 5, CartID: 5, RateBook: 7, RateBorrower: 7, DescRateBook: "asal", DescRateBorrower: "gtddnj"},
		{ID: 6, CartID: 6, RateBook: 6, RateBorrower: 7, DescRateBook: "asal", DescRateBorrower: "gtddnj"},
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
	db.Create(&Cart)
	db.Create(&Rating)

}
