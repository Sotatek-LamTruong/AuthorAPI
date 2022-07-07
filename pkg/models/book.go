package models

type Book struct {
	Idbook     int    `json:"id"`
	Bookname   string `json:"name"`
	AuthorId   int    `json:"author_id"`
	CategoryId int    `json:"category_id"`
}
