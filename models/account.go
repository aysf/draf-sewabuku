package models

import "time"

type Account struct {
	ID           uint `gorm:"primaryKey"`
	Balance      uint
	CreatedAt    time.Time
	TransferTo   []Transfers `gorm:"foreignKey:ToAccountId;references:ID"`
	TransferFrom []Transfers `gorm:"foreignKey:FromAccountId;references:ID"`
	UserID       uint
	User         User
}
