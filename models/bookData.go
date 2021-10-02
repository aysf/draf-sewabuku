package models

type BookData struct {
	ID          uint      `db:"id"`
	Tittle      string    `db:"tittle"`
	AuthorID    uint      `db:"author_id"`
	PublishYear uint      `db:"publish_year" json:"publish_year"`
	CategoryID  uint      `db:"category_id" json:"category_id"`
	PublisherID uint      `db:"publisher_id" json:"publisher_id"`
	Author      Author    `db:"authors" json:"authors"`
	Publisher   Publisher `db:"publishers" json:"publishers"`
	Category    Category  `db:"categories" json:"categories"`
}

type Publisher struct {
	ID   uint   `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name"`
}

type Author struct {
	ID   uint   `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name"`
}

type Category struct {
	ID   uint   `db:"id" json:"id,omitempty"`
	Name string `db:"name" json:"name"`
}

type BookDataResponse struct {
	ID        uint      `db:"id" json:"id"`
	Tittle    string    `db:"tittle" json:"tittle"`
	UserID    uint      `db:"user_id" json:"owner_id"`
	OwnerName string    `gorm:"owner_name" json:"owner_name"`
	Address   string    `json:"adress"`
	RentPrice uint16    `db:"rent_price" json:"rent_price"`
	Category  Category  `gorm:"categories" db:"categories" json:"category"`
	Publisher Publisher `gorm:"publishers" db:"publishers" json:"publishers"`
	Author    Author    `gorm:"authors" db:"authors" json:"authors"`
}

type InputBook struct {
	Title       string `json:"title"`
	CategoryID  uint   `json:"category_id"`
	AuthorID    uint   `json:"author_id"`
	PublisherID uint   `json:"publisher_id"`
	PublishYear uint   `json:"publish_year"`
	Price       uint16 `json:"price"`
}
