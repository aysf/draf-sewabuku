package database

import (
	"gorm.io/gorm"
	"sewabuku/models"
)

type (
	GormBookModel struct {
		db *gorm.DB
	}
	BookModel interface {
		GetAll() ([]models.Book, error)
		Get(bookId int) (models.Book, error)
		Insert(models.Book) (models.Book, error)
		Edit(book models.Book, bookId int) (models.Book, error)
		Delete(bookId int) (models.Book, error)
	}
)

func (g GormBookModel) GetAll() ([]models.Book, error) {
	panic("implement me")
}

func (g GormBookModel) Get(bookId int) (models.Book, error) {
	panic("implement me")
}

func (g GormBookModel) Insert(book models.Book) (models.Book, error) {
	panic("implement me")
}

func (g GormBookModel) Edit(book models.Book, bookId int) (models.Book, error) {
	panic("implement me")
}

func (g GormBookModel) Delete(bookId int) (models.Book, error) {
	panic("implement me")
}

func NewBookModel(db *gorm.DB) *GormBookModel {
	return &GormBookModel{db: db}
}
