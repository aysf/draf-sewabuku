package database

import (
<<<<<<< HEAD
	"sewabuku/middlewares"
	"sewabuku/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
=======
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"sewabuku/middlewares"
	"sewabuku/models"
	"strconv"
>>>>>>> 162fabe65615ad0dcbb4468ed551a5c7ed315a4f
)

type (
	GormUserModel struct {
		db *gorm.DB
	}
	UserModel interface {
		Register(user models.User) (models.User, error)
		Login(email, password string) (models.User, error)
		GetProfile(userId int) (models.User, error)
		UpdatePassword(newPass models.User, userId int) (models.User, error)
	}
)

// NewUserModel is function to initialize new user model
func NewUserModel(db *gorm.DB) *GormUserModel {
	db.Exec(`
	CREATE TRIGGER after_create_user
	AFTER INSERT ON users FOR EACH ROW 
	INSERT INTO accounts(balance, user_id)
	VALUES (0, new.id);`)
	return &GormUserModel{db: db}
}

// Register is  method to add new user
func (g *GormUserModel) Register(user models.User) (models.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		err := errors.New("ALL FIELD CANNOT EMPTY")
		return user, err
	}

	bcryptCost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	passwordEncrypted, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)

	user.Password = string(passwordEncrypted)

	if err := g.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Login is method to user log in
func (g *GormUserModel) Login(email, password string) (models.User, error) {
	var user models.User
	var err error

	if err = g.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return user, err
	}

	user.Token, _ = middlewares.CreateToken(int(user.ID))

	if err = g.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// GetProfile is  method to get user profile
func (g *GormUserModel) GetProfile(userId int) (models.User, error) {
	var user models.User

	if err := g.db.Find(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}

// UpdatePassword is method to edit user password
func (g *GormUserModel) UpdatePassword(newPass models.User, userId int) (models.User, error) {
	var user models.User
	var err error

	g.db.First(&user, userId)

	user.Password = newPass.Password

	if err = g.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
