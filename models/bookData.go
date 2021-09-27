package models

import "time"

type BookData struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	AuthorID      uint
	PublisherID   uint
	CategoryID    uint
	PublisherDate time.Time
	Author        Author
	Publisher     Publisher
	Category      Catagory
}

type Publisher struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Author struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Catagory struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
