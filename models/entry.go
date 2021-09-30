package models

import "time"

type Entry struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint
	Amount    int
	CreatedAt time.Time
	Account   Account
}
