package dto

import "book-author/pkg/models"

type AddBookReq struct {
	BookId     int    `json:"id"`
	BookName   string `json:"name"`
	AuthorId   int    `json:"author_id"`
	CategoryId int    `json:"category_id"`
}

// type addBookRes struct {
// 	Book *models.Book
// }

type GetBookByAuthorReq struct {
	AuthorId int `json:"auth_id"`
}

type GetBookByCateReq struct {
	CateId int `json:"cate_id"`
}

type GetBookByNameReq struct {
	BookName string `json:"name"`
}

type GetBookRes struct {
	Books []models.Book
}

type GetBookByNameRes struct {
	Book *models.Book
}
