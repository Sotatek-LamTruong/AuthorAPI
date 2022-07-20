package dto

import (
	"book-author/pkg/models"
)

type ListAuthor struct {
	Authors []AuthorDTO `json:"authors"`
}

type AuthorDTO struct {
	AuthorId   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
	// Category   models.Category `json:"category,omitempty"`
	Books []models.Book `json:"books,omitempty"`
}

type GetAuthorReq struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
}

type GetAuthorRes struct {
	AuthorId   int           `json:"author"`
	AuthorName string        `json:"author_name"`
	Books      []models.Book `json:"books"`
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

type GetAuthorsByBook struct {
	BookId   int             `json:"book_id"`
	BookName string          `json:"book_name"`
	Authors  []models.Author `json:"authors"`
}

type GetAuthorByBook struct {
	BookId     int             `json:"book_id"`
	BookName   string          `json:"book_name"`
	CategoryId int             `json:"category_id"`
	Authors    []models.Author `json:"authors"`
}
