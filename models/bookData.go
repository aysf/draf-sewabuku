package models

import "time"

type BookData struct {
	ID            uint `gorm:"primaryKey"`
	OwnerID       uint
	Title         string
	CategoryID    uint
	PublisherDate time.Time
	Author        string
	Publisher     string
	PublishDate   time.Time
	FileName      string
	PeiceBook     uint
}
type Catagory struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type InputBook struct {
	Title      string
	CategoryID uint
	Author     string
	Publisher  string
	Price      uint16
}
