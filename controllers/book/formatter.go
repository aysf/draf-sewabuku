package book

import "sewabuku/models"

func FormatResponseBook(input models.BookData) Formatter {
	format := Formatter{
		Available: input.Quantity > 0,
		ID:        input.ID,
		Title:     input.Title,
		Photo:     input.Photo,
		Price:     input.Price,
		Address:   input.User.Address,
		Author:    input.Author,
		Publisher: input.Publisher,
		Category:  input.Category,
	}

	return format

}

func FormatResponseBooks(books []models.BookData) []Formatter {
	formatter := []Formatter{}
	for _, v := range books {
		formatt := FormatResponseBook(v)
		formatter = append(formatter, formatt)
	}

	return formatter
}

func FormatDetailsBook(input models.BookData) FormatDetails {

	formatter := FormatDetails{
		ID:          input.ID,
		Title:       input.Title,
		Photo:       input.Photo,
		PublishYear: input.PublishYear,
		Quantity:    input.Quantity,
		Price:       input.Price,
		Rating:      input.Rating,
		Description: input.Description,
		Name:        input.User.Name,
		Address:     input.User.Address,
		Author:      input.Author,
		Publisher:   input.Publisher,
		Category:    input.Category,
	}
	return formatter
}

type Formatter struct {
	ID        uint             `db:"id" json:"id"`
	Title     string           `db:"Title" json:"Tittle"`
	Photo     string           `json:"file_foto,omitempty"`
	Price     uint             `json:"rent_price"`
	Address   string           `json:"address"`
	Author    models.Author    `db:"authors" json:"author"`
	Publisher models.Publisher `db:"publishers" json:"publisher"`
	Category  models.Category  `db:"categories" json:"category"`
	Available bool             `json:"is_available"`
}

type FormatDetails struct {
	ID          uint             `json:"id"`
	Title       string           `json:"Tittle"`
	Photo       string           `json:"file_foto"`
	PublishYear uint             `json:"publish_year"`
	Quantity    uint             `json:"quantity"`
	Price       uint             `json:"rent_price"`
	Rating      float32          `json:"rating"`
	Description string           `json:"description"`
	Name        string           `json:"owner_name"`
	Address     string           `json:"address"`
	Author      models.Author    `json:"author"`
	Publisher   models.Publisher `json:"publisher"`
	Category    models.Category  `json:"category"`
}
