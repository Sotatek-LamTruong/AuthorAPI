package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"fmt"
)

type AuthorServices interface {
	GetAllAuthors() (*dto.ListAuthor, error)
	GetAuthor(id int) (*dto.GetAuthorRes, error)
	CreateAuthor(dto.CreateAuthorReq) error
	GetAuthorByBook(bookID int) (*dto.GetAuthorByBookRes, error)
}

type DefaultAuthor struct {
	repo repository.AuthorRepository
}

func NewAuthor(repo repository.AuthorRepository) AuthorServices {
	return DefaultAuthor{
		repo: repo,
	}
}

func (d DefaultAuthor) GetAllAuthors() (*dto.ListAuthor, error) {
	authors, err := d.repo.List()
	if err != nil {
		fmt.Println(err)
	}
	for _, author := range authors {
		fmt.Println(author)
	}
	return &dto.ListAuthor{
		Authors: authors,
	}, nil
}

func (d DefaultAuthor) GetAuthor(id int) (*dto.GetAuthorRes, error) {
	author, err := d.repo.Get(id)

	if err != nil {
		fmt.Println(err)
	}

	return &dto.GetAuthorRes{Author: author}, nil

}

func (d DefaultAuthor) CreateAuthor(author dto.CreateAuthorReq) error {
	result := models.Author{
		Name: author.Name,
	}

	err := d.repo.Create(&result)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (d DefaultAuthor) GetAuthorByBook(bookID int) (*dto.GetAuthorByBookRes, error) {
	author, err := d.repo.GetByBook(bookID)

	if err != nil {
		fmt.Println(err)
	}

	return &dto.GetAuthorByBookRes{Author: author}, nil
}
