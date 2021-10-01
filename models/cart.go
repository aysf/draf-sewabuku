package models

import "time"

type Cart struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	BookUserID uint
	DateLoan   time.Time
	DateDue    time.Time
	DateReturn time.Time
	User       User
	BookUser   BookUser
}
