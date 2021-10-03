package models

type Book struct {
	ID          uint   `db:"id"`
	Title       string `db:"title"`
	AuthorID    uint   `db:"author_id" json:"author_id"`
	PublishYear uint   `db:"publish_year" json:"publish_year"`
	CategoryID  uint   `db:"category_id" json:"category_id"`
	PublisherID uint   `db:"publisher_id" json:"publisher_id"`
	UserID      uint
	RentPrice   uint
	Quantity    uint
	Rating      uint
	Description string
	FileFoto    string
	Author      Author    `db:"authors" json:"authors"`
	Publisher   Publisher `db:"publishers" json:"publishers"`
	Category    Category  `db:"categories" json:"categories"`
}

type BookData struct {
	ID          uint      `db:"id"`
	Tittle      string    `db:"tittle"`
	AuthorID    uint      `db:"author_id" json:"author_id"`
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

type BookRespone struct {
	ID          uint      `db:"id" json:"id"`
	Tittle      string    `db:"tittle" json:"tittle"`
	Photo       string    `json:"file_foto,omitempty"`
	PublishYear uint      `db:"publish_year" json:"publish_year,omitempty"`
	UserID      uint      `db:"users_id" json:"owner_id,omitempty"`
	OwnerName   string    `gorm:"owner_name" json:"owner_name,omitempty"`
	Quantity    uint      `json:"quantity"`
	Address     string    `json:"address"`
	Description string    `json:"description,omitempty"`
	PublisherID uint      `db:"publisher_id" json:"-"`
	AuthorID    uint      `db:"author_id" json:"-"`
	CategoryID  uint      `db:"category_id" json:"-"`
	Author      Author    `db:"authors" json:"authors"`
	Publisher   Publisher `db:"publishers" json:"publishers"`
	Category    Category  `db:"categories" json:"categories"`
}

type InputBook struct {
	Title       string `json:"title"`
	CategoryID  uint   `json:"category_id"`
	AuthorID    uint   `json:"author_id"`
	PublisherID uint   `json:"publisher_id"`
	PublishYear uint   `json:"publish_year"`
	Price       uint16 `json:"price"`
	Quantity    uint   `json:"quantity"`
	Description string `jon:"description"`
}
