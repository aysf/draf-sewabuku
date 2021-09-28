package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email" gorm:"unique"`
	Password  string `json:"-" form:"password"`
	Token     string `json:"token"`
	AccountID uint
	Account   Account
	CreatedAt time.Time `json:"-"`
}
