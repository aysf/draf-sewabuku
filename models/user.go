package models

import "time"

type User struct {
<<<<<<< HEAD
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"-" form:"password"`
	Token    string `json:"token"`
	// AccountID uint
	// Account   Account
=======
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name" form:"name" gorm:"not null"`
	Email     string `json:"email" form:"email" gorm:"unique;not null"`
	Password  string `json:"-" form:"password" gorm:"not null"`
	Token     string `json:"token"`
	//AccountID uint
	//Account   Account
>>>>>>> 162fabe65615ad0dcbb4468ed551a5c7ed315a4f
	CreatedAt time.Time `json:"-"`
}
