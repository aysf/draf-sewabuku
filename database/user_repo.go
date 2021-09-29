package database

import (
	"fmt"
	"sewabuku/middlewares"
	"sewabuku/models"

	"gorm.io/gorm"
)

type (
	GormUserModel struct {
		db *gorm.DB
	}
	UserModel interface {
		Register(user models.User) (models.User, error)
		Login(email, password string) (models.User, error)
		GetProfile(userId int) (models.User, error)
	}
)

func NewUserModel(db *gorm.DB) *GormUserModel {
	if err := db.Exec(`
	CREATE TRIGGER after_create_user
	AFTER INSERT ON users FOR EACH ROW 
	INSERT INTO accounts(balance, user_id)
	VALUES (0, new.id)`); err != nil {
		fmt.Println("error")
	}
	return &GormUserModel{db: db}
}

func (g GormUserModel) Register(user models.User) (models.User, error) {
	if err := g.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (g GormUserModel) Login(email, password string) (models.User, error) {
	var user models.User
	var err error

	if err = g.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))

	if err != nil {
		return user, err
	}

	if err = g.db.Save(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (g GormUserModel) GetProfile(userId int) (models.User, error) {
	var user models.User

	if err := g.db.Find(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}
