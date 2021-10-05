package models

import "time"

type User struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name" form:"name" gorm:"not null" validate:"required"`
	OrganizationName string    `json:"organization_name" form:"organization_name"`
	Email            string    `json:"email" form:"email" gorm:"unique;not null" validate:"required,email"`
	Password         string    `json:"password" form:"password" gorm:"not null" validate:"required"`
	Address          string    `json:"address" form:"address" gorm:"default:none"`
	Token            string    `json:"token,omitempty"`
	CreatedAt        time.Time `json:"-"`
}
