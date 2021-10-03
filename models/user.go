package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" form:"name" gorm:"not null"`
	Email     string    `json:"email" form:"email" gorm:"unique;not null"`
	Password  string    `json:"password" form:"password" gorm:"not null"`
	Address   string    `json:"address" form:"address" gorm:"default:none"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"-"`
}

type UserResponse struct {
	Name    string `json:"name" form:"name" gorm:"not null"`
	Address string `json:"address" form:"address" gorm:"default:none"`
}
