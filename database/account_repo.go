package database

import (
	"errors"
	"sewabuku/models"

	"gorm.io/gorm"
)

type (
	GormAccountModel struct {
		db *gorm.DB
	}
	AccountModel interface {
		Show(userId int) (models.Account, error)
		Transaction(entry models.Entry) (models.Entry, error)
		UpdateBalance(id, amount uint) (interface{}, error)
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

func (g GormAccountModel) Transaction(entry models.Entry) (models.Entry, error) {
	if err := g.db.Create(&entry).Error; err != nil {
		return entry, err
	}
	return entry, nil
}

func (g GormAccountModel) UpdateBalance(id, amount uint) (interface{}, error) {
	var account models.Account
	var accountHold models.AccountHold

	if err := g.db.Where("user_id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	if err := g.db.Where("account_id = ?", account.ID).First(&accountHold).Error; err != nil {
		return nil, err
	}

	if account.Balance < amount {
		err := errors.New("your account balance is not enough to make the deposit")
		return nil, err
	}

	updateBalance := account.Balance - amount
	updateDeposit := accountHold.Balance + amount

	if err := g.db.Model(&account).Where("user_id = ?", id).Update("balance", updateBalance).Error; err != nil {
		return nil, err
	}
	if err := g.db.Model(&accountHold).Where("account_id = ?", account.ID).Update("balance", updateDeposit).Error; err != nil {
		return nil, err
	}

	return amount, nil

}
