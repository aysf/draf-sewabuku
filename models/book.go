package models

type Book struct {
	ID          uint `gorm:"primaryKey"`
	BookDataID  uint
	UserID      uint
	RentPrice   uint
	Quantity    uint
	Description string
	BookData    BookData
	User        User
}
