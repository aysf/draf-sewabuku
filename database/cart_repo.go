package database

import (
	"errors"
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
		GetBookByUserId(userId int) ([]models.BookData, error)
		Return(Date time.Time, userId, bookId int) (interface{}, error)
		List(userId int) ([]models.Cart, error)
		Extend(Date time.Time, userId, bookId int) (interface{}, error)
	}
)

// Rent is method to get book loan registration number
func (g *GormCartModel) Rent(cart models.Cart) (models.Cart, error) {

	if err := g.db.Create(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (g *GormCartModel) GetBookByUserId(userId int) ([]models.BookData, error) {
	var bookList []models.BookData

	if err := g.db.Model(&models.BookData{}).Where("user_id = ?", userId).Find(&bookList).Error; err != nil {
		return bookList, err
	}
	return bookList, nil
}

// Return is methor to update return book date
func (g *GormCartModel) Return(Date time.Time, userId, bookId int) (interface{}, error) {

	type Cart struct {
		ID         uint
		BookDataID uint
		UserID     uint
		DateDue    time.Time
		DateReturn time.Time
	}

	var cart Cart

	tx := g.db.Where("user_id = ? AND book_data_id = ?", userId, bookId).Last(&cart)
	if tx.Error != nil {
		return cart, tx.Error
	}
	var nullTime time.Time
	if cart.DateReturn != nullTime {
		return cart, errors.New("book already returned")
	}

	if err := g.db.Model(&cart).Update("date_return", Date).Error; err != nil {
		return cart, err
	}

	return cart, nil

}

// Return is methor to update return book date
func (g *GormCartModel) List(userId int) ([]models.Cart, error) {

	var carts []models.Cart
	if err := g.db.Model(&models.Cart{}).Where("user_id = ?", userId).Find(&carts).Error; err != nil {
		return carts, err
	}

	return carts, nil
}

func (g *GormCartModel) Extend(inputDate time.Time, userId, bookId int) (interface{}, error) {
	type Cart struct {
		ID         uint
		BookDataID uint
		DateDue    time.Time
	}

	var cart Cart

	tx := g.db.Model(&models.Cart{}).Where("user_id = ? AND book_data_id = ?", userId, bookId).Last(&cart)
	if tx.Error != nil {
		fmt.Println("cek2")
		return cart, tx.Error
	}
	today := time.Now()
	if today.Before(cart.DateDue) {
		fmt.Println("cek3")
		return nil, errors.New("the due date already expired, could not extend the loan period")
	}

	if inputDate.Before(cart.DateDue) {
		fmt.Println("cek4")
		fmt.Println(inputDate)
		fmt.Println(cart.DateDue)
		return nil, errors.New("could not input date before due, please correct your input")
	}

	cart.DateDue = inputDate

	if err := g.db.Save(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

// NewCartModel is function to initialize new cart model
func NewCartModel(db *gorm.DB) *GormCartModel {

	db.Exec(`
	CREATE TRIGGER after_cart_insert_lender
	AFTER INSERT ON carts
	FOR EACH ROW
		INSERT INTO entries (account_id, amount, created_at) 
		VALUES ((select user_id from book_data where book_data.id = new.book_data_id), DATEDIFF(new.date_due, new.date_loan) * (select price from book_data where book_data.id = new.book_data_id), now()); `)

	db.Exec(`
	CREATE TRIGGER after_cart_insert_borrower
	AFTER INSERT ON carts
	FOR EACH ROW
		INSERT INTO entries (account_id, amount, created_at) 
		VALUES (new.user_id, DATEDIFF(  new.date_due, new.date_loan) *(select -1*CAST(price AS SIGNED) from book_data where book_data.id = new.book_data_id), now());`)

	db.Exec(`
	CREATE TRIGGER after_rent_qty
	AFTER INSERT on carts
	for each row 
		update book_data set quantity = quantity - 1
		where book_data.id = new.book_data_id;
	`)

	db.Exec(`
	CREATE TRIGGER after_return_qty
	AFTER UPDATE on carts
	for each row 
		IF new.date_return <> old.date_return THEN
				update book_data set quantity = quantity + 1
				where book_data.id = new.book_data_id;
		END IF
	`)

	db.Exec(`
	CREATE TRIGGER after_return_ratings
	AFTER UPDATE on carts
	for each row 
		IF new.date_return <> old.date_return THEN
				INSERT INTO ratings(cart_id) VALUES(new.id);
		END IF
	`)

	db.Exec(`
	CREATE TRIGGER after_extend_balance_lender
	AFTER UPDATE on carts
	for each row 
		IF new.date_due <> old.date_due THEN
			INSERT INTO entries (account_id, amount, created_at) 
			VALUES ((select user_id from book_data where book_data.id = new.book_data_id), DATEDIFF(new.date_due, new.date_loan) * (select price from book_data where book_data.id = new.book_data_id), now());
		END IF
	`)

	db.Exec(`
	CREATE TRIGGER after_extend_balance_borrower
	AFTER UPDATE on carts
	for each row 
		IF new.date_due <> old.date_due THEN
		INSERT INTO entries (account_id, amount, created_at) 
		VALUES (new.user_id, DATEDIFF(  new.date_due, new.date_loan) *(select -1*CAST(price AS SIGNED) from book_data where book_data.id = new.book_data_id), now());
		END IF
	`)

	return &GormCartModel{db: db}
}
