package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" form:"name" gorm:"not null"`
	Email    string `json:"email" form:"email" gorm:"unique;not null"`
	Password string `json:"-" form:"password" gorm:"not null"`
	Token    string `json:"token"`
	//AccountID uint
	//Account   Account
	CreatedAt time.Time `json:"-"`
}
