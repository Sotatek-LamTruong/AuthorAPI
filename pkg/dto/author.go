package dto

import (
	"book-author/pkg/models"
)

type ListAuthor struct {
	Authors []*models.Author `json:"authors"`
}

//
type GetAuthorReq struct {
	AuthorID int `json:"author_id"`
}

type GetAuthorRes struct {
	Author *models.Author `json:"author"`
}

//
type CreateAuthorReq struct {
	AuthorID int    `json:"author_id"`
	Name     string `json:"name"`
}

type CreateAuthorRes struct {
	Author *models.Author `json:"author"`
}

//

type GetAuthorByBookReq struct {
	Book *models.Book `json:"book"`
	// BookId int          `json:"book_id"`
}

type GetAuthorByBookRes struct {
	Author *models.Author `json:"author"`
}
