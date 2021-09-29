package models

import "time"

type BookData struct {
	ID            uint `gorm:"primaryKey"`
	OwnerID       uint
	Title         string
	CategoryID    uint
	Author        string
	Publisher     string
	PublishDate   time.Time
	PhotoFileName string
	PeiceBook     uint16
}
type Catagory struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type InputBook struct {
	Title         string    `json:"tittle"`
	CategoryID    uint      `json:"category_id"`
	Author        string    `json:"author"`
	Publisher     string    `json:"publisher"`
	PublishDate   time.Time `json:"publish_date"`
	PhotoFileName string    `json:"photo_file"`
	Price         uint16    `json:"price"`
}
