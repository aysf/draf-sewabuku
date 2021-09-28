package models

import "time"

type Account struct {
	ID        uint `gorm:"primaryKey"`
	Balance   uint
	CreatedAt time.Time
	// User         User        `gorm:"foreignKey:AccountId;references:ID"`
	TransferTo   []Transfers `gorm:"foreignKey:ToAccountId;references:ID"`
	TransferFrom []Transfers `gorm:"foreignKey:FromAccountId;references:ID"`
}
