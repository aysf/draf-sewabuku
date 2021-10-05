package database

import (
	"sewabuku/models"
	"time"

	"gorm.io/gorm"
)

type (
	GormCartModel struct {
		db *gorm.DB
	}
	CartModel interface {
		Rent(cart models.Cart) (models.Cart, error)
		Return(Date time.Time, userId, bookId int) (models.Cart, error)
		List(userId int) (interface{}, error)
	}
)

// Rent is method to get book loan registration number
func (g *GormCartModel) Rent(cart models.Cart) (models.Cart, error) {

	if err := g.db.Create(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

// Return is methor to update return book date
func (g *GormCartModel) Return(Date time.Time, userId, bookId int) (models.Cart, error) {
	var cart models.Cart

	tx := g.db.Where("user_id = ? AND book_user_id = ?", userId, bookId).Find(&cart)
	if tx.Error != nil {
		return cart, tx.Error
	}

	if err := g.db.Model(&cart).Update("date_return", Date).Error; err != nil {
		return cart, err
	}

	return cart, nil

}

// Return is methor to update return book date
func (g *GormCartModel) List(userId int) (interface{}, error) {

	type CartView struct {
		ID         uint
		BookUserID uint
		DateLoan   time.Time
		DateDue    time.Time
		DateReturn time.Time
		IsReturn   bool
	}

	var carts []CartView

	if err := g.db.Model(&models.Cart{}).Where("user_id = ?", userId).Find(&carts).Error; err != nil {
		return carts, err
	}

	return carts, nil
}

// NewCartModel is function to initialize new cart model
func NewCartModel(db *gorm.DB) *GormCartModel {

	db.Exec(`
	CREATE TRIGGER after_cart_insert_lender
	AFTER INSERT ON carts
	FOR EACH ROW
	INSERT INTO entries (account_id, amount, created_at) VALUES ((select user_id from book_data where book_data.id = new.book_data_id), DATEDIFF(new.date_due, new.date_loan) * (select price from book_data where book_data.id = new.book_data_id), now()); `)

	db.Exec(`
	CREATE TRIGGER after_cart_insert_borrower
	AFTER INSERT ON carts
	FOR EACH ROW
	INSERT INTO entries (account_id, amount, created_at) VALUES (new.user_id, DATEDIFF(  new.date_due, new.date_loan) *(select -1*CAST(price AS SIGNED) from book_data where book_data.id = new.book_data_id), now());`)
	return &GormCartModel{db: db}
}
