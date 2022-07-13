package models

type Book struct {
	IdBook     int    `json:"id"`
	BookName   string `json:"name_book"`
	CategoryId int    `json:"category_id,omitempty"`
}
