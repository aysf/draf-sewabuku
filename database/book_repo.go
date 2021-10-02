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
		GetAllBooks() ([]models.BookDataResponse, error)
		GetByCategoryID(id int) ([]models.BookData, error)
		GetAll() ([]models.BookData, error)
		Search(keyword, author, category string) (interface{}, error)
		Get(bookId int) (models.BookUser, error)
		GetByNameBook(namebook string) ([]models.BookData, error)
		ListAuthor() ([]models.Author, error)
		InputBook(input models.BookData) (models.BookData, error)
		GetListPublisher() ([]models.Publisher, error)
		Insert(models.BookUser) (models.BookUser, error)
		Edit(book models.BookUser, bookId int) (models.BookUser, error)
		Delete(bookId int) (models.BookUser, error)
		UpdateBook(input models.BookData) (models.BookData, error)
		GetBookByID(id uint) (models.BookData, error)
		GetByAuthorID(id int) ([]models.BookData, error)
		GetByPublisherID(id int) ([]models.BookData, error)
		ListCategory() ([]models.Category, error)
		CreateNewAuthor(input models.Author) (models.Author, error)
		CreateNewPublisher(input models.Publisher) (models.Publisher, error)
	}
)

func (r *GormBookModel) GetAllBooks() ([]models.BookDataResponse, error) {
	var books []models.BookDataResponse

	querry := `SELECT book_data.id, tittle, publish_year, users.id, users.address,book_users.rent_price,book_users.file_foto,users.name, authors.name, authors.id,publishers.name, publishers.id,categories.name, categories.id
	FROM book_data
	JOIN authors ON authors.id = book_data.author_id
	JOIN publishers ON publishers.id = book_data.publisher_id
	JOIN categories ON categories.id = category_id
	JOIN book_users ON book_data.id = book_users.book_data_id
	JOIN users ON book_users.user_id = users.id`
	err := r.db.Preload("Authors").Preload("Publishers").Preload("Categories").Raw(querry).Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, err
}

func (r *GormBookModel) GetByCategoryID(id int) ([]models.BookData, error) {
	var books []models.BookData

	querry := `SELECT book_data.id, tittle, publish_year, users.id, users.address,book_users.rent_price,book_users.file_foto,users.name, authors.name, authors.id,publishers.name, publishers.id,categories.name, categories.id
	FROM book_data
	JOIN authors ON authors.id = book_data.author_id
	JOIN publishers ON publishers.id = book_data.publisher_id
	JOIN categories ON categories.id = category_id
	JOIN book_users ON book_data.id = book_users.book_data_id
	JOIN users ON book_users.user_id = users.id WHERE categories.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Raw(querry, id).Find(&books).Error
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

	querry := `SELECT bd.*, c.id as "categories.id", c.name as "categories.name", p.id as "publishers.id", p.name as "publishers.name", a.id as "authors.id", a.name as "authors.name" FROM book_data bd JOIN categories c ON bd.category_id = c.id JOIN publishers p ON bd.publisher_id = p.id JOIN authors a ON bd.author_id = a.id WHERE bd.title LIKE ?`
	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Raw(querry, "%"+namebook+"%").Find(&books).Error
	if err != nil {
		return []models.BookData{}, err
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

func (r *GormBookModel) UpdateBook(input models.BookData) (models.BookData, error) {

	err := r.db.Save(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *GormBookModel) GetBookByID(id uint) (models.BookData, error) {
	var book models.BookData

	querry := `SELECT bd.*, c.id as "categories.id", c.name as "categories.name", p.id as "publishers.id", p.name as "publishers.name", a.id as "authors.id", a.name as "authors.name" FROM book_data bd JOIN categories c ON bd.category_id = c.id JOIN publishers p ON bd.publisher_id = p.id JOIN authors a ON bd.author_id = a.id WHERE bd.id = ?`
	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Raw(querry, id).Find(&book).Error

	if err != nil {
		return models.BookData{}, err
	}

	return book, nil
}

func (r *GormBookModel) GetByAuthorID(id int) ([]models.BookData, error) {
	var books []models.BookData
	querry := `SELECT bd.*, c.id as "categories.id", c.name as "categories.name", p.id as "publishers.id", p.name as "publishers.name", a.id as "authors.id", a.name as "authors.name" FROM book_data bd JOIN categories c ON bd.category_id = c.id JOIN publishers p ON bd.publisher_id = p.id JOIN authors a ON bd.author_id = a.id WHERE a.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Raw(querry, id).Find(&books).Error

	if err != nil {
		return []models.BookData{}, err
	}

	return books, nil
}

func (r *GormBookModel) GetByPublisherID(id int) ([]models.BookData, error) {
	var books []models.BookData
	querry := `SELECT bd.*, c.id as "categories.id", c.name as "categories.name", p.id as "publishers.id", p.name as "publishers.name", a.id as "authors.id", a.name as "authors.name" FROM book_data bd JOIN categories c ON bd.category_id = c.id JOIN publishers p ON bd.publisher_id = p.id JOIN authors a ON bd.author_id = a.id WHERE p.id = ?`

	err := r.db.Preload("Author").Preload("Publisher").Preload("Category").Raw(querry, id).Find(&books).Error

	if err != nil {
		return []models.BookData{}, err
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
	err := db.Exec(`CREATE OR REPLACE VIEW responsebook AS
	SELECT book_data.id as id, tittle, publish_year, users.id AS user_id, users.address as address,book_users.rent_price as rent_price,book_users.file_foto as photo,users.name AS owner_name, authors.name AS author_name, authors.id as author_id ,publishers.name as publisher_name, publishers.id as publisher_id ,categories.name AS category_name, categories.id as category_id
		FROM book_data
		LEFT JOIN authors ON authors.id = book_data.author_id
		LEFT JOIN publishers ON publishers.id = book_data.publisher_id
		LEFT JOIN categories ON categories.id = category_id
		LEFT JOIN book_users ON book_data.id = book_users.book_data_id
		LEFT JOIN users ON book_users.user_id = users.id`).Error
	if err != nil {
		fmt.Println(err)
		panic("panic")
	}

	return &GormBookModel{db: db}
}
