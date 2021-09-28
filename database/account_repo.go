package database

import (
	"sewabuku/models"

	"gorm.io/gorm"
)

type (
	GormAccountModel struct {
		db *gorm.DB
	}
	AccountModel interface {
		Show(userId int) (models.Account, error)
		Add(bookId int) (models.Account, error)
	}
)

func NewAccountModel(db *gorm.DB) *GormAccountModel {
	return &GormAccountModel{db: db}
}

func (g GormAccountModel) Show(userId int) (models.Account, error) {
	var account models.Account

	if err := g.db.Find(&account, userId).Error; err != nil {
		return account, err
	}

	return account, nil
}

func (g GormAccountModel) Add(account models.Account) (models.Account, error) {
	if err := g.db.Create(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}
