package models

import "time"

type Cart struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	BookDataID uint      `json:"book_user_id" form:"book_user_id" gorm:"not null"`
	DateLoan   time.Time `json:"date_loan" form:"date_loan" gorm:"not null"`
	DateDue    time.Time `json:"date_due" form:"date_due" gorm:"not null"`
	DateReturn time.Time `json:"date_return" form:"date_return"`
	IsReturn   bool
	User       User
	BookData   BookData
}
