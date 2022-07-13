package dto

import "book-author/pkg/models"

type CreateCateReq struct {
	CategoryID   string `json:"id"`
	CategoryName string `json:"name"`
}

type CreateCateRes struct {
	Category *models.Category `json:"category"`
}

type GetCateByBookReq struct {
	IDBook int `json:"book_id"`
}

type GetCateRes struct {
	CategoryId   int           `json:"id"`
	CategoryName string        `json:"name"`
	Books        []models.Book `json:"books"`
}

type GetCateByNameReq struct {
	Name string `json:"name"`
}
