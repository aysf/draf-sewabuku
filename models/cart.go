package models

import "time"

type Cart struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	BookDataID uint      `json:"book_user_id" form:"book_user_id" gorm:"not null"`
	DateLoan   time.Time `json:"date_loan" form:"date_loan" gorm:"not null"`
	DateDue    time.Time `json:"date_due" form:"date_due" gorm:"not null"`
	DateReturn time.Time `json:"date_return" form:"date_return"`
	User       User
	BookData   BookData
}

type Rating struct {
	ID               uint `gorm:"primaryKey"`
	CartID           uint
	RateBook         float32 `json:"rate_book" form:"rate_book" gorm:"not null"`
	RateBorrower     float32 `json:"rate_borrower" form:"rate_borrower" gorm:"not null"`
	DescRateBook     string  `json:"desc_rate_book" form:"desc_rate_book" gorm:"not null"`
	DescRateBorrower string  `json:"desc_rate_borrower" form:"desc_rate_borrower" gorm:"not null"`
	Cart             Cart
}
