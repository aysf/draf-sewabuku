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
		{Name: "Ami", Email: "ami@mail.com", Password: passStr, Address: "jakarta"},
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
		{AccountID: "a-1", Amount: 25000, CreatedAt: time.Now()},
		{AccountID: "a-2", Amount: 50000, CreatedAt: time.Now()},
		{AccountID: "a-3", Amount: 75000, CreatedAt: time.Now()},
		{AccountID: "a-4", Amount: 100000, CreatedAt: time.Now()},
	}

	// ----------------
	// Book Data
	// ----------------

	var bookData = []models.BookData{
		{
			Title:       "Rich Dad Poor Dad", //1
			UserID:      1,
			Quantity:    1,
			Photo:       "default.jpg",
			AuthorID:    3,
			PublisherID: 2,
			CategoryID:  4,
			PublishYear: 1997,
			Price:       100,
		},
		{
			Title:       "Kambing Jantan", //2
			UserID:      1,
			Photo:       "default.jpg",
			Quantity:    1,
			AuthorID:    4,
			PublisherID: 4,
			CategoryID:  5,
			PublishYear: 2002,
			Price:       100,
		},
		{
			Title:       "Rumah Seribu Malaikat", //3
			UserID:      8,
			Photo:       "default.jpg",
			Quantity:    1,
			AuthorID:    5,
			PublisherID: 3,
			CategoryID:  5,
			PublishYear: 2009,
			Price:       100,
		},
		{
			Title:       "The Cruel Prince", //4
			UserID:      7,
			Photo:       ".jpg",
			Quantity:    0,
			AuthorID:    6,
			PublisherID: 5,
			CategoryID:  6,
			PublishYear: 2009,
			Price:       100,
		},
		{
			Title:       "The Black Box", //5
			UserID:      6,
			Quantity:    1,
			Photo:       ".jpg",
			AuthorID:    7,
			PublisherID: 6,
			CategoryID:  3,
			PublishYear: 2009,
			Price:       200,
		},
		{
			Title:       "Black Clover", //6
			UserID:      5,
			AuthorID:    2,
			Quantity:    1,
			Photo:       ".jpg",
			PublisherID: 1,
			CategoryID:  6,
			PublishYear: 1999,
			Price:       100,
		},
		{
			Title:       "La Tahzan", //7
			UserID:      4,
			Quantity:    1,
			Photo:       "jpg",
			AuthorID:    11,
			PublisherID: 3,
			CategoryID:  3,
			PublishYear: 2009,
			Price:       10,
		},
		{
			Title:       "Sang Pemimpi", //8
			UserID:      3,
			Quantity:    1,
			Photo:       "default.jpg",
			AuthorID:    2,
			PublisherID: 9,
			CategoryID:  2,
			PublishYear: 2010,
			Price:       100,
		}, {
			Title:       "Juru Masak Para Maiko 05", //9
			UserID:      2,
			Quantity:    2,
			Photo:       "default.jpg",
			AuthorID:    1,
			PublisherID: 2,
			CategoryID:  3,
			PublishYear: 2010,
			Price:       100,
		}, {
			Title:       "Teknik Pemograman Pascal", //10
			UserID:      1,
			Quantity:    5,
			Photo:       "default.jpg",
			AuthorID:    9,
			PublisherID: 2,
			CategoryID:  5,
			PublishYear: 2010,
			Price:       100,
		}, {
			Title:       "Hacking is Easy", //11
			UserID:      1,
			Quantity:    1,
			Photo:       "default.jpg",
			AuthorID:    10,
			PublisherID: 8,
			CategoryID:  5,
			PublishYear: 2010,
			Price:       400,
		},
	}

	// ----------------
	// Author
	// ----------------

	var Author = []models.Author{
		{Name: "none"},              //1
		{Name: "Andrea Hirata"},     //2
		{Name: "Robert Kiyosaki"},   //3
		{Name: "Raditya Dika"},      //4
		{Name: "Yuli Badawi"},       //5
		{Name: "Holly Black"},       //6
		{Name: "Michael, Connelly"}, //7
		{Name: "Tabata & Yuuki"},    //8
		{Name: "Budi Raharjo"},      //9
		{Name: "Efvy Zam Kerinci"},  //10
		{Name: "Aid al-Qarni"},      //11
	}

	// ----------------
	// Category
	// ----------------

	var Category = []models.Category{
		{Name: "none"},      //1
		{Name: "sejarah"},   //2
		{Name: "novel"},     //3
		{Name: "motivasi"},  //4
		{Name: "non fiksi"}, //5
		{Name: "comic"},     //6
		{Name: "Agama"},     //7
		{Name: "horror"},    //8
	}

	// ----------------
	// Publisher
	// ----------------

	var Publisher = []models.Publisher{
		{Name: "none"},                 //1
		{Name: "gramedia"},             //2
		{Name: "mizan"},                //3
		{Name: "andi"},                 //4
		{Name: "LBYR"},                 //5
		{Name: "Orion"},                //6
		{Name: "Penerbit Informatika"}, //7
		{Name: "Neomedia press"},       //8
		{Name: "Bentang Pustaka"},      //9
	}

	// ----------------
	// Cart and Rating mock
	// ----------------

	var Cart = []models.Cart{
		// rate case 1 - for case book id = 10 got ratings from 4 borrowers
		{ID: 1, UserID: 5, BookDataID: 10, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		{ID: 2, UserID: 6, BookDataID: 10, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		{ID: 3, UserID: 7, BookDataID: 10, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		{ID: 4, UserID: 8, BookDataID: 10, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		// for case user id = 5 got ratings from 4 lenders
		{ID: 5, UserID: 5, BookDataID: 7, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		{ID: 6, UserID: 5, BookDataID: 8, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		{ID: 7, UserID: 5, BookDataID: 9, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
		{ID: 8, UserID: 5, BookDataID: 10, DateLoan: time.Date(2020, 11, 1, 0, 0, 0, 0, time.UTC), DateDue: time.Date(2020, 11, 5, 0, 0, 0, 0, time.UTC), DateReturn: time.Date(2020, 11, 4, 0, 0, 0, 0, time.UTC), IsReturn: true},
	}

	var Rating = []models.Rating{
		// rate case 1
		{ID: 1, CartID: 1, RateBook: 8, RateBorrower: 6, DescRateBook: "bukunya bagus", DescRateBorrower: "mengembalikan tepat waktu"},
		{ID: 2, CartID: 2, RateBook: 5, RateBorrower: 8, DescRateBook: "beberapa halaman ada yang kotor", DescRateBorrower: "mengembalikan buku tepat waktu"},
		{ID: 3, CartID: 3, RateBook: 7, RateBorrower: 9, DescRateBook: "sangat bermanfaat", DescRateBorrower: "peminjam sangat merawat bukunya"},
		{ID: 4, CartID: 4, RateBook: 8, RateBorrower: 6, DescRateBook: "bukunya keren", DescRateBorrower: "beberapa halaman lecek/lusuh"},
		// rate case 2
		{ID: 5, CartID: 5, RateBook: 7, RateBorrower: 9, DescRateBook: "bukunya keren bingitz", DescRateBorrower: "bukunya masih terjaga"},
		{ID: 6, CartID: 6, RateBook: 9, RateBorrower: 8, DescRateBook: "asyik bacanya, pengen pinjem lagi kapan2", DescRateBorrower: "bukunya dikembalikan tepat waktu"},
		{ID: 7, CartID: 7, RateBook: 8, RateBorrower: 5, DescRateBook: "recomended book, usernya juga ramah", DescRateBorrower: "bukunya dikembalikan dalam keadaan basah"},
		{ID: 8, CartID: 8, RateBook: 6, RateBorrower: 6, DescRateBook: "isi bukunya terlalu sulit dipahami", DescRateBorrower: "sedikit terlambat"},
	}

	db.Create(&User)
	db.Create(&entry)
	db.Create(&Category)
	db.Create(&Author)
	db.Create(&Publisher)
	db.Create(&bookData)
	db.Create(&Cart)
	db.Create(&Rating)

}
