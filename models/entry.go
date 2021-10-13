package models

import "time"

type Entry struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	AccountID string    `json:"-"`
	Amount    int       `json:"amount" form:"amount"`
	CreatedAt time.Time `json:"-"`
	Account   Account   `json:"-"`
}
