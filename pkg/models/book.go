package models

type Book struct {
	Idbook   int      `json:"id"`
	Bookname string   `json:"name"`
	Author   Author   `json:"author"`
	Category Category `json:"category"`
}
