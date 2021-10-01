package models

type BookData struct {
	ID            uint `gorm:"primaryKey"`
	Title         string
	AuthorID      uint `gorm:"default:1"`
	PublisherID   uint `gorm:"default:1"`
	CategoryID    uint `gorm:"default:1"`
	PublisherYear uint ``
	Author        Author
	Publisher     Publisher
	Category      Category
}

type Publisher struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type AuthorBook struct {
}

type Author struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
