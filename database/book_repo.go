package database

import (
	"sewabuku/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type RepositoryBook interface {
	GetBookByID(id uint) (models.BookData, error)
	GetAllBooks() ([]models.BookData, error)
	GetByCategory(categoryname string) ([]models.BookData, error)
	GetByNameBook(namebook string) ([]models.BookData, error)
	InputBook(input models.BookData) (models.BookData, error)
	UpdateBook(input models.BookData) (models.BookData, error)
	GetByAuthor(name string) ([]models.BookData, error)
	GetByPublisher(name string) ([]models.BookData, error)
	DeleteBook(id uint) error
	ListCategory() ([]models.Catagory, error)
}

func NewBookRepostory(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllBooks() ([]models.BookData, error) {
	var books []models.BookData
	err := r.db.Raw("SELECT * FROM book_data").Scan(&books).Error

	if err != nil {
		return []models.BookData{}, err
	}

	return books, err
}

func (r *repository) GetByCategory(categoryname string) ([]models.BookData, error) {
	var books []models.BookData

	err := r.db.Raw("SELECT * FROM book_data INNER JOIN catagories ON book_data.category_id = catagories.id WHERE catagories.name = ?", categoryname).Scan(&books).Error
	if err != nil {
		return []models.BookData{}, err
	}

	return books, nil
}

func (r *repository) GetByNameBook(namebook string) ([]models.BookData, error) {
	var books []models.BookData

	err := r.db.Where("title LIKE ?", "%"+namebook+"%").Find(&books).Error
	if err != nil {
		return []models.BookData{}, err
	}

	return books, nil
}

func (r *repository) InputBook(input models.BookData) (models.BookData, error) {

	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) UpdateBook(input models.BookData) (models.BookData, error) {

	err := r.db.Save(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}

func (r *repository) GetBookByID(id uint) (models.BookData, error) {
	var book models.BookData

	err := r.db.Raw("SELECT * FROM book_data WHERE id = ?", id).Scan(&book).Error

	if err != nil {
		return models.BookData{}, err
	}

	return book, nil
}

func (r *repository) GetByAuthor(name string) ([]models.BookData, error) {
	var books []models.BookData

	err := r.db.Raw("SELECT * FROM book_data WHERE author = ?", name).Find(&books).Error

	if err != nil {
		return []models.BookData{}, err
	}

	return books, nil
}

func (r *repository) GetByPublisher(name string) ([]models.BookData, error) {
	var books []models.BookData

	err := r.db.Raw("SELECT * FROM book_data WHERE publisher = ?", name).Find(&books).Error

	if err != nil {
		return []models.BookData{}, err
	}

	return books, nil
}

func (r *repository) DeleteBook(id uint) error {

	var book models.BookData
	err := r.db.Where("id = ?", id).Delete(&book).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ListCategory() ([]models.Catagory, error) {
	var category []models.Catagory

	err := r.db.Raw("SELECT * FROM catagories").Find(&category).Error
	if err != nil {
		return []models.Catagory{}, err
	}

	return category, nil
}
