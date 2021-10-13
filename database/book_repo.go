package database

import (
	"fmt"
	"sewabuku/models"

	"gorm.io/gorm"
)

type (
	GormBookModel struct {
		db *gorm.DB
	}
	BookModel interface {
		GetAllBooks() ([]models.BookData, error)
		ListAuthor() ([]models.Author, error)
		InputBook(input models.BookData) (models.BookData, error)
		GetListPublisher() ([]models.Publisher, error)
		GetBookByID(id uint) (models.BookData, error)
		GetByKeywordID(author, category, publisher int) ([]models.BookData, error)
		ListCategory() ([]models.Category, error)
		CreateNewAuthor(input models.Author) (models.Author, error)
		CreateNewPublisher(input models.Publisher) (models.Publisher, error)
		UpdatePhoto(file string, book_id int) (models.BookData, error)
		CheckAuthorName(name string) (bool, error)
		CheckPublisherName(name string) (bool, error)
		BorrowBook(cart models.Cart) (models.Cart, error)
		InsertNewBook(input models.BookData) (models.BookData, error)
		CheckBorrowBook(user_id int) (bool, error)
		SearchBooks(keyword string, author, publisher, category int) ([]models.BookData, error)
	}
)

func (r *GormBookModel) GetAllBooks() ([]models.BookData, error) {
	var books []models.BookData
	querry := `SELECT b.id, b.title, b.photo, b.publish_year, b.price, b.quantity, b.description, b.user_id, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id`
	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (g *GormBookModel) SearchBooks(keyword string, author, publisher, category int) ([]models.BookData, error) {

	var books []models.BookData
	var tx *gorm.DB

	if author != 0 && category != 0 && publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND author_id = ? AND category_id = ? AND publisher_id = ?", keyword, author, category, publisher).Find(&books)
	} else if category != 0 && publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND category_id = ? AND publisher_id = ?", keyword, category, publisher).Find(&books)
	} else if author != 0 && category != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND author_id = ? AND category_id = ?", keyword, author, category).Find(&books)
	} else if author != 0 && publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND author_id = ? AND publisher_id = ?", keyword, author, publisher).Find(&books)
	} else if author != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND author_id = ?", keyword, author).Find(&books)
	} else if publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND publisher_id = ?", keyword, publisher).Find(&books)
	} else if category != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ? AND category_id = ?", keyword, category).Find(&books)
	} else {
		tx = g.db.Preload("User").Preload("Author").Preload("Publisher").Preload("Category").Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS ('', title, author, category,publisher) LIKE ?", keyword).Find(&books)
	}

	if tx.Error != nil {
		return books, tx.Error
	}

	return books, nil

}

func (r *GormBookModel) ListAuthor() ([]models.Author, error) {
	var authors []models.Author

	querry := `SELECT * FROM authors`
	err := r.db.Raw(querry).Find(&authors).Error
	if err != nil {
		return []models.Author{}, err
	}

	return authors, nil
}

func (r *GormBookModel) GetListPublisher() ([]models.Publisher, error) {
	var publishers []models.Publisher
	querry := `SELECT * FROM publishers`

	err := r.db.Raw(querry).Find(&publishers).Error
	if err != nil {
		return []models.Publisher{}, err
	}

	return publishers, nil

}

func (r *GormBookModel) InputBook(input models.BookData) (models.BookData, error) {

	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *GormBookModel) GetBookByID(id uint) (models.BookData, error) {
	var book models.BookData

	querry := `SELECT b.*, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id WHERE b.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry, id).Find(&book).Error
	if err != nil {
		return book, err
	}

	querry2 := `SELECT AVG(rate_book) rate FROM rating r WHERE book_data_id = ?`
	var rate float32
	err = r.db.Raw(querry2, id).Find(&rate).Error
	if err != nil {
		return book, nil
	}
	book.Rating = rate

	return book, nil
}

func (g *GormBookModel) GetByKeywordID(author, category, publisher int) ([]models.BookData, error) {
	var books []models.BookData
	var tx *gorm.DB

	if author != 0 && category != 0 && publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE author_id = ? AND category_id = ? AND publisher_id = ?", author, category, publisher).Find(&books)
	} else if category != 0 && publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE category_id = ? AND publisher_id = ?", category, publisher).Find(&books)
	} else if author != 0 && category != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE author_id = ? AND category_id = ?", author, category).Find(&books)
	} else if author != 0 && publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE author_id = ? AND publisher_id = ?", author, publisher).Find(&books)
	} else if author != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE author_id = ?", author).Find(&books)
	} else if publisher != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE publisher_id = ?", publisher).Find(&books)
	} else if category != 0 {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs WHERE category_id = ?", category).Find(&books)
	} else {
		tx = g.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw("SELECT * FROM book_catalogs").Find(&books)
	}

	if tx.Error != nil {
		return books, tx.Error
	}

	return books, nil
}

func (r *GormBookModel) DeleteBook(id uint) error {

	var book models.BookData
	err := r.db.Where("id = ?", id).Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *GormBookModel) ListCategory() ([]models.Category, error) {
	var category []models.Category
	querry := `SELECT * FROM categories`

	err := r.db.Raw(querry).Find(&category).Error
	if err != nil {
		return []models.Category{}, err
	}

	return category, nil
}

func (r *GormBookModel) CreateNewAuthor(input models.Author) (models.Author, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return models.Author{}, err
	}

	return input, nil
}

func (r *GormBookModel) CreateNewPublisher(input models.Publisher) (models.Publisher, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return models.Publisher{}, err
	}

	return input, nil
}

func (r *GormBookModel) UpdatePhoto(file string, book_id int) (models.BookData, error) {

	var response models.BookData
	err := r.db.Model(&models.BookData{}).Where("id", book_id).Update("photo", file).Error
	if err != nil {
		return response, err
	}

	return response, nil
}

func (r *GormBookModel) CheckAuthorName(name string) (bool, error) {

	var author models.Author

	querry := `SELECT * FROM authors WHERE name = ?`

	err := r.db.Raw(querry, name).Find(&author).Error
	if err != nil {
		return false, err
	}
	if author.Name != "" {
		return false, nil
	}

	return true, nil
}

func (r *GormBookModel) CheckPublisherName(name string) (bool, error) {

	var publisher models.Publisher

	querry := `SELECT * FROM publishers WHERE name = ?`

	err := r.db.Raw(querry, name).Find(&publisher).Error
	if err != nil {
		return false, err
	}
	if publisher.Name != "" {
		return false, nil
	}

	return true, nil
}

func (r *GormBookModel) InsertNewBook(input models.BookData) (models.BookData, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	querry := `SELECT b.*, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id WHERE b.id = ?`

	var book models.BookData
	err = r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry, input.ID).Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *GormBookModel) CheckBorrowBook(user_id int) (bool, error) {
	qurry := `SELECT * FROM carts WHERE user_id = ?`

	var cart models.Cart
	err := r.db.Raw(qurry, user_id).Find(&cart).Error
	if err != nil {
		return false, err
	}
	if cart.BookDataID != 0 {
		return false, nil
	}

	return true, nil
}

func (r *GormBookModel) BorrowBook(cart models.Cart) (models.Cart, error) {

	err := r.db.Create(&cart).Error
	if err != nil {
		return cart, err
	}

	err = r.db.Model(&models.BookData{}).Where("id", cart.BookDataID).Update("quantity = ?", "quantity - 1").Error
	if err != nil {
		return cart, err
	}

	return cart, nil

}

func NewBookModel(db *gorm.DB) *GormBookModel {
	err := db.Exec(`CREATE OR REPLACE VIEW book_catalogs AS
	SELECT 
		bd.id as id, 
		bd.title as title, 
		bd.description as description , 
		bd.price as price, 
		bd.photo as photo, 
		bd.quantity as quantity , 
		bd.publish_year as publish_year, 
		bd.publisher_id as publisher_id , 
		bd.author_id as author_id , 
		bd.category_id as category_id ,
		a.name AS author, 
		p.name as publisher, 
		c.name as category, 
		u.id as user_id, 
		u.address as address, 
		u.name as name
	FROM book_data bd 
	LEFT JOIN authors a ON a.id = bd.author_id
	LEFT JOIN publishers p ON p.id = bd.publisher_id
	LEFT JOIN categories c ON c.id = bd.category_id
	LEFT JOIN  users u ON bd.user_id = u.id`).Error
	if err != nil {
		fmt.Println("masih errorrrrr disni")
		panic(err)
	}

	err1 := db.Exec(`CREATE OR REPLACE VIEW rating AS
	SELECT 
		b.id AS book_data_id, 
		r.rate_book AS rate_book 
	FROM book_data b
	LEFT JOIN carts c on b.id = c.book_data_id
	LEFT JOIN ratings r on c.id = r.cart_id`).Error

	if err1 != nil {
		fmt.Println("masih errorrrrr disni")
		panic(err)
	}

	return &GormBookModel{db: db}
}
