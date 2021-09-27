package database

import (
	"gorm.io/gorm"
	"sewabuku/models"
)

type (
	GormUserModel struct {
		db *gorm.DB
	}
	UserModel interface {
		GetAll() ([]models.User, error)
		Get(userId int) (models.User, error)
		Insert(models.User) (models.User, error)
		Edit(user models.User, userId int) (models.User, error)
		Delete(userId int) (models.User, error)
		Login(email, password string) (models.User, error)
	}
)

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}

func (g GormUserModel) GetAll() ([]models.User, error) {
	panic("implement me")
}

func (g GormUserModel) Get(userId int) (models.User, error) {
	panic("implement me")
}

func (g GormUserModel) Insert(user models.User) (models.User, error) {
	panic("implement me")
}

func (g GormUserModel) Edit(user models.User, userId int) (models.User, error) {
	panic("implement me")
}

func (g GormUserModel) Delete(userId int) (models.User, error) {
	panic("implement me")
}

func (g GormUserModel) Login(email, password string) (models.User, error) {
	panic("implement me")
}
