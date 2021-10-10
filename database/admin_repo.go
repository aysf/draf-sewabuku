package database

import (
	"gorm.io/gorm"
	"sewabuku/models"
)

type (
	GormAdminModel struct {
		db *gorm.DB
	}
	AdminModel interface {
		GetAllUser() (models.User, error)
		GetUserHistory(userId int) (models.User, error)
	}
)

// NewAdminModel is function to initialize new user model
func NewAdminModel(db *gorm.DB) *GormAdminModel {
	return &GormAdminModel{db: db}
}