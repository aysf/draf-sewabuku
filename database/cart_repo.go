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
		GetAccountByUserId(userId int) (models.Account, models.AccountHold, error)
		GetLenderIdByBookId(bookId uint) (uint, error)
		GetBookByBookId(bookId uint) (models.BookData, error)
		UpdateSaldo(account models.Account, amount int) (interface{}, error)
		Return(Date time.Time, userId, bookId int) (interface{}, error)
		List(userId int) ([]models.Cart, error)
		Extend(Date time.Time, userId, bookId int) (interface{}, error)
	}
)

// get lender id
func (g *GormCartModel) GetLenderIdByBookId(bookId uint) (uint, error) {
	type Result struct {
		UserID uint
	}
	var result Result

	if err := g.db.Table("book_data").Select("user_id").Where("id = ?", bookId).Scan(&result).Error; err != nil {
		return result.UserID, err
	}
	// if err := g.db.Raw("SELECT user_id FROM book_data WHERE id = ?", bookId).Scan(&user).Error; err != nil {
	// 	return user.ID, err
	// }

	return result.UserID, nil
}

// UpdateSaldo
func (g *GormCartModel) UpdateSaldo(account models.Account, amount int) (interface{}, error) {

	updateBalance := account.Balance + uint(amount)

	if err := g.db.Model(&account).Where("id = ?", account.ID).Update("balance", updateBalance).Error; err != nil {
		return nil, err
	}
	return updateBalance, nil
}

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

func (g *GormCartModel) GetAccountByUserId(userId int) (models.Account, models.AccountHold, error) {
	var account models.Account
	var accountHold models.AccountHold

	if err := g.db.Model(&models.Account{}).Where("user_id = ?", userId).Find(&account).Error; err != nil {
		return account, accountHold, err
	}

	if err := g.db.Model(&models.AccountHold{}).Where("account_id = ?", account.ID).Find(&accountHold).Error; err != nil {
		return account, accountHold, err
	}

	return account, accountHold, nil
}

func (g *GormCartModel) GetBookByBookId(bookId uint) (models.BookData, error) {
	var book models.BookData
	if err := g.db.Model(&models.BookData{}).Where("id = ?", bookId).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
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
		msg := errors.New("input invalid")
		return cart, msg
	}
	var nullTime time.Time
	if cart.DateReturn != nullTime {
		return cart, errors.New("book already returned")
	}

	if err := g.db.Model(&cart).Update("date_return", Date).Error; err != nil {
		msg := errors.New("update failed")
		return cart, msg
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

	//added trigger transaction
	/*
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
	*/

	// db.Exec(`
	// CREATE TRIGGER after_transfer_rent_lender
	// AFTER INSERT on transfers
	// FOR EACH ROW
	// 	INSERT INTO entries (account_id, amount, created_at)
	// 	VALUES
	// 		(new.to_account_id,
	// 		new.amount,
	// 		now());
	// `)

	/*

		db.Exec(`
		CREATE TRIGGER after_transfer_rent_lender
		AFTER INSERT on transfers
		FOR EACH ROW
		UPDATE accounts
			SET balance = balance + new.amount, updated_at = now()
			WHERE id = new.to_account_id;
		`)

		db.Exec(`
		CREATE TRIGGER after_transfer_rent_borrower
		AFTER INSERT on transfers
		FOR EACH ROW
		UPDATE accounts
			SET balance = balance - new.amount, updated_at = now()
			WHERE id = new.from_account_id;
		`)
	*/

	// update balance account after deposit or withdrawal in entries table
	db.Exec(`
	CREATE TRIGGER after_entries_insert
	AFTER INSERT ON entries 
	FOR EACH ROW
		UPDATE accounts
		SET balance = balance + new.amount, updated_at = now()
		WHERE id = new.account_id;`)

	// --------------------------------------------
	// update table transfer after rent transaction
	// --------------------------------------------

	db.Exec(`
	create trigger after_cart_rent_transfer
	after insert on carts
	for each row
	insert into transfers (to_account_id, from_account_id, amount, created_at)
	values 
		((select a.id from accounts a where a.user_id = (select bd.user_id from book_data bd where bd.id = new.book_data_id)), 
		(select a.id from accounts a where a.user_id = new.user_id), 
		DATEDIFF(new.date_due, new.date_loan) * (select price from book_data where book_data.id = new.book_data_id), 
		now());
	`)

	// ------------------------------------------------------
	// update balance in table account after rent transaction
	// ------------------------------------------------------

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

	// ------------------------------------------------------
	// create ratings record after book return
	// ------------------------------------------------------

	db.Exec(`
	CREATE TRIGGER after_return_ratings
	AFTER UPDATE on carts
	for each row 
		IF new.date_return <> old.date_return THEN
				INSERT INTO ratings(cart_id) VALUES(new.id);
		END IF
	`)

	// ------------------------------------------------------
	// add transaction for extending book rental period
	// ------------------------------------------------------

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
