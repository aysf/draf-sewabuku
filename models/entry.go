package models

import "time"

type Entry struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint
	Amount    int `json:"amount" form:"amount"`
	CreatedAt time.Time
	Account   Account
}
