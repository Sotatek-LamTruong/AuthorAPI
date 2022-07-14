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

//
type GetAuthorReq struct {
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
}

type GetAuthorRes struct {
	AuthorId   int    `json:"author"`
	AuthorName string `json:"author_name"`
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

type GetAuthorByBookRes struct {
	Author *models.Author `json:"author"`
}
