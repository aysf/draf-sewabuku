package database

import (
	"fmt"
	"sewabuku/models"

	"gorm.io/gorm"
)

type (
	GormAccountModel struct {
		db *gorm.DB
	}
	AccountModel interface {
		Show(userId int) (models.Account, error)
		Add(entry models.Entry) (models.Entry, error)
	}
)

func NewAccountModel(db *gorm.DB) *GormAccountModel {
	if err := db.Exec(`
	CREATE TRIGGER after_entries_insert 
	AFTER INSERT ON entries FOR EACH ROW 
	UPDATE accounts SET balance = balance + new.amount`); err != nil {
		fmt.Println("error")
	}
	return &GormAccountModel{db: db}

}

func (g GormAccountModel) Show(userId int) (models.Account, error) {
	var account models.Account

	if err := g.db.Find(&account, userId).Error; err != nil {
		return account, err
	}

	return account, nil
}

func (g GormAccountModel) Add(entry models.Entry) (models.Entry, error) {
	if err := g.db.Create(&entry).Error; err != nil {
		return entry, err
	}

	return entry, nil
}
