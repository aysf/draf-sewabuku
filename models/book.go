package models

type Book struct {
	ID         uint `gorm:"primaryKey"`
	BookDataID uint
	UserID     uint
	BookData   BookData
	User       User
}
