package database

import (
	"fmt"
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
	fmt.Println("row yg terpengaruh -->", tx.RowsAffected)
	fmt.Println("cek format data -->", cart.DateReturn)

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
	INSERT INTO entries (account_id, amount) VALUES ((select user_id from book_users where book_users.id = new.book_user_id), DATEDIFF(new.date_due, new.date_loan) * (select rent_price from book_users where book_users.id = new.book_user_id)); `)

	db.Exec(`
	CREATE TRIGGER after_cart_insert_borrower
	AFTER INSERT ON carts
	FOR EACH ROW
	INSERT INTO entries (account_id, amount) VALUES (new.user_id, DATEDIFF(  new.date_due, new.date_loan) *(select -1*CAST(rent_price AS SIGNED) from book_users where book_users.id = new.book_user_id));`)
	return &GormCartModel{db: db}
}
