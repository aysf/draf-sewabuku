package database

import (
	"fmt"
	"sewabuku/models"
	"time"

	"gorm.io/gorm"
)

type (
	GormBookModel struct {
		db *gorm.DB
	}
	BookModel interface {
		GetAllBooks() ([]models.BookData, error)
		GetByCategoryID(id int) ([]models.BookData, error)
		Search(keyword, author, category string) (interface{}, error)
		GetByNameBook(namebook string) ([]models.BookData, error)
		ListAuthor() ([]models.Author, error)
		InputBook(input models.BookData) (models.BookData, error)
		GetListPublisher() ([]models.Publisher, error)
		GetBookByID(id uint) (models.BookData, error)
		GetByAuthorID(id int) ([]models.BookData, error)
		GetByPublisherID(id int) ([]models.BookData, error)
		ListCategory() ([]models.Category, error)
		CreateNewAuthor(input models.Author) (models.Author, error)
		CreateNewPublisher(input models.Publisher) (models.Publisher, error)
		UpdatePhoto(file string, book_id int) (models.BookData, error)
		CheckAuthorName(name string) (bool, error)
		CheckPublisherName(name string) (bool, error)
		BorrowBook(book_id, user_id int) (models.Cart, error)
		InsertNewBook(input models.BookData) (models.BookData, error)
		CheckBorrowBook(user_id int) (bool, error)
	}
)

func (r *GormBookModel) GetAllBooks() ([]models.BookData, error) {
	var books []models.BookData

	querry := `SELECT b.id, b.tittle, b.photo, b.publish_year, b.price, b.quantity, b.description, b.user_id, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *GormBookModel) GetByCategoryID(id int) ([]models.BookData, error) {
	var books []models.BookData

	querry := `SELECT b.*, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id WHERE c.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry, id).Find(&books).Error
	if err != nil {
		return books, err
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

func (r *GormBookModel) GetByNameBook(namebook string) ([]models.BookData, error) {
	var books []models.BookData

	querry := `SELECT b.id, b.tittle, b.photo, b.publish_year, b.price, b.quantity, b.description, b.user_id, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id WHERE b.tittle LIKE ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry, "%"+namebook+"%").Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
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

	return book, nil
}

func (r *GormBookModel) GetByAuthorID(id int) ([]models.BookData, error) {
	var books []models.BookData

	querry := `SELECT b.*, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id WHERE a.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry, id).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *GormBookModel) GetByPublisherID(id int) ([]models.BookData, error) {
	var books []models.BookData

	querry := `SELECT b.*, u.address as "users.address", u.name as "users.name", a.name as "authors.name" ,b.author_id, p.name as "publishers.name", publisher_id, c.name as "categories.name", category_id FROM book_data b JOIN users u ON b.user_id = u.id JOIN publishers p ON b.publisher_id = p.id JOIN authors a ON b.author_id = a.id JOIN categories c ON b.category_id = c.id WHERE p.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Preload("User").Raw(querry, id).Find(&books).Error
	if err != nil {
		return books, err
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

func (r *GormBookModel) BorrowBook(book_id, user_id int) (models.Cart, error) {
	cart := models.Cart{
		UserID:     uint(user_id),
		BookDataID: uint(book_id),
		DateLoan:   time.Now(),
		DateDue:    <-time.After(time.Hour * 240),
		DateReturn: <-time.After(time.Hour * 240),
	}

	err := r.db.Create(&cart).Error
	if err != nil {
		return cart, err
	}

	err = r.db.Model(&models.BookData{}).Where("id", book_id).Update("quantity = ?", -1).Error
	if err != nil {
		return cart, err
	}

	return cart, nil

}

func (g GormBookModel) Search(keyword, author, category string) (interface{}, error) {
	type BookCatalog struct {
		Title       string
		PublishYear uint
		Author      string
		Publisher   string
		Category    string
	}

	var result []BookCatalog

	var tx *gorm.DB

	if author != "%%" && category != "%%" {
		tx = g.db.Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS('', title, author, publisher, category) LIKE ? AND author LIKE ? AND category LIKE ?", keyword, author, category).Scan(&result)
	} else if category != "%%" {
		tx = g.db.Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS('', title, author, publisher, category) LIKE ? AND category LIKE ?", keyword, category).Scan(&result)
	} else if author != "%%" {
		tx = g.db.Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS('', title, author, publisher, category) LIKE ? AND author LIKE ?", keyword, author).Scan(&result)
	} else {
		tx = g.db.Raw("SELECT * FROM book_catalogs WHERE CONCAT_WS('', title, author, publisher, category) LIKE ?", keyword).Scan(&result)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func NewBookModel(db *gorm.DB) *GormBookModel {
	if err := db.Exec(`CREATE VIEW book_catalogs AS
	SELECT title, publish_year, authors.name AS author, publishers.name as publisher, categories.name as category 
	FROM book_data
	LEFT JOIN authors ON authors.id = book_data.author_id
	LEFT JOIN publishers ON publishers.id = book_data.publisher_id
	LEFT JOIN categories ON categories.id = category_id`); err != nil {
		fmt.Println("there is error during loading trigger after_entries_insert")
	}

	return &GormBookModel{db: db}
}
