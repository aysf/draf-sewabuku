package models

type BookData struct {
	ID          uint      `db:"id" json:"id"`
	Title       string    `db:"title" json:"tittle"`
	Photo       string    `json:"file_foto,omitempty"`
	PublishYear uint      `db:"publish_year" json:"publish_year,omitempty"`
	Quantity    uint      `json:"quantity"`
	Price       uint      `json:"rent_price"`
	Description string    `json:"description,omitempty"`
	Rating      float32   `json:"rating"`
	UserID      uint      `db:"user_id" json:"-"`
	User        User      `db:"users" json:"user"`
	PublisherID uint      `db:"publisher_id" json:"-"`
	AuthorID    uint      `db:"author_id" json:"-"`
	CategoryID  uint      `db:"category_id" json:"-"`
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

type InputBook struct {
	Title       string `json:"Title"`
	CategoryID  uint   `json:"category_id"`
	AuthorID    uint   `json:"author_id"`
	PublisherID uint   `json:"publisher_id"`
	PublishYear uint   `json:"publish_year"`
	Price       uint16 `json:"price"`
	Quantity    uint   `json:"quantity"`
	Description string `jon:"description"`
}
