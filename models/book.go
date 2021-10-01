package models

type BookUser struct {
	ID          uint `gorm:"primaryKey"`
	BookDataID  uint
	UserID      uint
	RentPrice   uint
	Quantity    uint
	Rating      uint
	Description string
	BookData    BookData
	User        User
}
