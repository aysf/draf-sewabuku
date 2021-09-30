package models

import "time"

type Account struct {
	ID           uint `gorm:"primaryKey"`
	Balance      uint
	TransferTo   []Transfers `gorm:"foreignKey:ToAccountId;references:ID"`
	TransferFrom []Transfers `gorm:"foreignKey:FromAccountId;references:ID"`
	UpdatedAt    time.Time
	UserID       uint
	User         User
}
