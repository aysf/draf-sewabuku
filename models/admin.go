package models

import "time"

type Admin struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name" form:"name" gorm:"not null" validate:"required"`
	Email            string    `json:"email" form:"email" gorm:"unique;not null" validate:"required,email"`
	Password         string    `json:"password" form:"password" gorm:"not null" validate:"required"`
	CreatedAt        time.Time `json:"-"`
}