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
		GetAll() ([]models.BookData, error)
		Search(keyword, author, category string) (interface{}, error)
		Get(bookId int) (models.BookUser, error)
		Insert(models.BookUser) (models.BookUser, error)
		Edit(book models.BookUser, bookId int) (models.BookUser, error)
		Delete(bookId int) (models.BookUser, error)
	}
)

func (g GormBookModel) GetAll() ([]models.BookData, error) {
	listBook := new([]models.BookData)
	if err := g.db.Find(&listBook).Error; err != nil {
		return nil, err
	}
	return *listBook, nil
}

func (g GormBookModel) Search(keyword, author, category string) (interface{}, error) {
	type BookCatalog struct {
		Title         string
		PublisherYear uint
		Author        string
		Publisher     string
		Category      string
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

func (g GormBookModel) AddBook(book models.BookUser) (models.BookUser, error) {
	panic("implement me")
}

func (g GormBookModel) Get(bookId int) (models.BookUser, error) {
	panic("implement me")
}

func (g GormBookModel) Insert(book models.BookUser) (models.BookUser, error) {
	panic("implement me")
}

func (g GormBookModel) Edit(book models.BookUser, bookId int) (models.BookUser, error) {
	panic("implement me")
}

func (g GormBookModel) Delete(bookId int) (models.BookUser, error) {
	panic("implement me")
}

func NewBookModel(db *gorm.DB) *GormBookModel {
	if err := db.Exec(`CREATE VIEW book_catalogs AS
	SELECT title, publisher_year, authors.name AS author, publishers.name as publisher, categories.name as category 
	FROM book_data
	LEFT JOIN authors ON authors.id = book_data.author_id
	LEFT JOIN publishers ON publishers.id = book_data.publisher_id
	LEFT JOIN categories ON categories.id = category_id`); err != nil {
		fmt.Println("there is error during loading trigger after_entries_insert")
	}

	return &GormBookModel{db: db}
}
