package dto

import "book-author/pkg/models"

type AddBookReq struct {
	BookId   int    `json:"id"`
	BookName string `json:"name"`
}

type CateReq struct {
	CategoryID int    `json:"category_id"`
	Name       string `json:"name,omitempty"`
}

type AuthorReq struct {
	AuthorID int    `json:"author_id"`
	Name     string `json:"name,omitempty"`
}

type AuthorRes struct {
	Row      int64  `json:"row"`
	BookID   int    `json:"book_id"`
	BookName string `json:"book_name"`
}

type GetByAuthor struct {
	AuthorID   int           `json:"author_id"`
	AuthorName string        `json:"author_name"`
	Books      []models.Book `json:"books"`
}

type GetByCate struct {
	CateID   int           `json:"category_id"`
	CateName string        `json:"category_name"`
	Books    []models.Book `json:"books"`
}

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

type UpdateAuthorByBookReq struct {
	Name string `json:"name"`
}
