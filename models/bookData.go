package models

import "time"

type BookData struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	OwnerID       uint      `json:"user_id"`
	Title         string    `json:"title"`
	CategoryID    uint      `json:"category_id"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	PublishDate   time.Time `json:"publish_date"`
	PhotoFileName string    `json:"photo_file"`
	PeiceBook     uint16    `json:"price"`
}
type Catagory struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Publisher struct {
	ID   uint
	Name string
}

type Author struct {
	ID   uint
	Name string
}

type InputBook struct {
	Title         string    `json:"title"`
	CategoryID    uint      `json:"category_id"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	PublishDate   time.Time `json:"publish_date"`
	PhotoFileName string    `json:"photo_file"`
	Price         uint16    `json:"price"`
}
