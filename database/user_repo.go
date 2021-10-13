package database

import (
	"os"
	"sewabuku/middlewares"
	"sewabuku/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	GormUserModel struct {
		db *gorm.DB
	}
	UserProfile struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Address string `json:"address"`
		Balance uint   `json:"balance"`
	}
	Borrowed []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
		PublishYear uint   `json:"publish_year"`
		Photo       string `json:"photo"`
		Price       uint   `json:"price"`
		Owner       string `json:"owner"`
		DateLoan    string `json:"date_loan"`
		DateDue     string `json:"date_due"`
		DateReturn  string `json:"date_return"`
	}
	Lent []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
		PublishYear uint   `json:"publish_year"`
		Photo       string `json:"photo"`
		Price       uint   `json:"price"`
		Borrower    string `json:"borrower"`
		DateLoan    string `json:"date_loan"`
		DateDue     string `json:"date_due"`
		DateReturn  string `json:"date_return"`
	}
	UserModel interface {
		Register(user models.User) (models.User, error)
		Login(email, password string) (models.User, error)
		GetProfile(userId int) (UserProfile, error)
		UpdateProfile(newProfile models.User, userId int) (models.User, error)
		UpdatePassword(newPass models.User, userId int) (models.User, error)
		Logout(userId int) (models.User, error)
		GetBorrowed(userId int, complete string) (Borrowed, error)
		GetLent(userId int, complete string) (Lent, error)
		InsertRating(rating models.Rating) (models.Rating, error)
		//InsertLentRating(userId int) (models.Rating, error)
	}
)

// NewUserModel is function to initialize new user model
func NewUserModel(db *gorm.DB) *GormUserModel {

	// create account and account_hold table
	db.Exec(`
		CREATE TRIGGER create_account
		AFTER INSERT ON users 
		FOR EACH ROW 
			INSERT INTO accounts(balance, user_id, id)
			VALUES (0, new.id, concat_ws('-',"a",new.id));`)

	db.Exec(`
		CREATE TRIGGER create_account_hold
		AFTER INSERT ON accounts 
		FOR EACH ROW 
			INSERT INTO account_holds(balance, account_id, id)
			VALUES (0, new.id, concat_ws('-',"d", (select users.id from users where new.user_id = users.id) ));`)

	db.Exec(`CREATE OR REPLACE VIEW user_profile AS
	SELECT 	users.id,
			users.name,
			users.organization_name,
			users.email,
			users.address,
        	accounts.balance
	FROM users
	LEFT JOIN accounts ON users.id = accounts.user_id;`)

	db.Exec(`CREATE OR REPLACE VIEW book_history AS
	SELECT book_data.user_id AS owner_id,
	       lender.name       AS owner,
	       title,
	       description,
	       categories.name   AS category,
	       authors.name      AS author,
	       publishers.name   AS publisher,
	       publish_year,
	       photo,
	       price,
	       carts.user_id     AS borrower_id,
	       borrower.name     AS borrower,
	       date_loan,
	       date_due,
	       date_return
	FROM book_data
	         LEFT JOIN carts ON book_data.id = carts.book_data_id
	         LEFT JOIN categories ON book_data.category_id = categories.id
	         LEFT JOIN authors ON book_data.author_id = authors.id
	         LEFT JOIN publishers ON book_data.publisher_id = publishers.id
	         LEFT JOIN users AS lender ON book_data.user_id = lender.id
	         LEFT JOIN users AS borrower ON carts.user_id = borrower.id
	WHERE carts.user_id IS NOT NULL;`)

	return &GormUserModel{db: db}
}

// Register is  method to add new user
func (g *GormUserModel) Register(user models.User) (models.User, error) {
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
func (g *GormUserModel) GetProfile(userId int) (UserProfile, error) {
	var user UserProfile

	if err := g.db.Raw("SELECT * FROM user_profile WHERE id = ?", userId).Scan(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// UpdateProfile is  method to edit user profile
func (g *GormUserModel) UpdateProfile(newProfile models.User, userId int) (models.User, error) {
	var user models.User
	var err error

	if err = g.db.First(&user, userId).Error; err != nil {
		return user, err
	}

	if newProfile.Name != "" {
		user.Name = newProfile.Name
	}

	if newProfile.OrganizationName != "" {
		user.OrganizationName = newProfile.OrganizationName
	}

	if newProfile.Email != "" {
		user.Email = newProfile.Email
	}

	if newProfile.Address != "" {
		user.Address = newProfile.Address
	}

	if err = g.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// UpdatePassword is method to edit user password
func (g *GormUserModel) UpdatePassword(newPass models.User, userId int) (models.User, error) {
	var user models.User
	var err error

	if err = g.db.First(&user, userId).Error; err != nil {
		return user, err
	}

	bcryptCost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	passwordEncrypted, _ := bcrypt.GenerateFromPassword([]byte(newPass.Password), bcryptCost)

	user.Password = string(passwordEncrypted)

	if err = g.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Logout is method to user log out
func (g *GormUserModel) Logout(userId int) (models.User, error) {
	var user models.User
	var err error

	if err = g.db.First(&user, userId).Error; err != nil {
		return user, err
	}

	user.Token = ""

	if err = g.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// GetBorrowed is method for get borrowed book
func (g *GormUserModel) GetBorrowed(userId int, complete string) (Borrowed, error) {
	var user Borrowed
	var err *gorm.DB

	if complete == "true" {
		err = g.db.Raw("SELECT * FROM book_history WHERE borrower_id = ? AND date_return IS NOT NULL", userId).Scan(&user)
	} else if complete == "false" {
		err = g.db.Raw("SELECT * FROM book_history WHERE borrower_id = ? AND date_return IS NULL", userId).Scan(&user)
	} else {
		err = g.db.Raw("SELECT * FROM book_history WHERE borrower_id = ?", userId).Scan(&user)
	}

	if err.Error != nil {
		return user, err.Error
	}

	return user, nil
}

// GetLent is method for get lent book
func (g *GormUserModel) GetLent(userId int, complete string) (Lent, error) {
	var user Lent
	var err *gorm.DB

	if complete == "true" {
		err = g.db.Raw("SELECT * FROM book_history WHERE owner_id = ? AND date_return IS NOT NULL", userId).Scan(&user)
	} else if complete == "false" {
		err = g.db.Raw("SELECT * FROM book_history WHERE owner_id = ? AND date_return IS NULL", userId).Scan(&user)
	} else {
		err = g.db.Raw("SELECT * FROM book_history WHERE owner_id = ?", userId).Scan(&user)
	}

	if err.Error != nil {
		return user, err.Error
	}

	return user, nil
}

// InsertRating is method for insert rating from borrower to lender
func (g *GormUserModel) InsertRating(rating models.Rating) (models.Rating, error) {
	if err := g.db.Save(&rating).Error; err != nil {
		return rating, err
	}

	return rating, nil
}

// InsertLentRating is method for insert rating from lender to borrower
//func (g *GormUserModel) InsertLentRating(userId int) (models.Rating, error) {
//
//}
