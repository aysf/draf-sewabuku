package models

import "time"

type LoanBook struct {
	ID         uint `gorm:"primaryKey"`
	BookDataID uint
	BookData   BookData `gorm:"foreignKey:BookDataID"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	MaxReturn  time.Time
}
