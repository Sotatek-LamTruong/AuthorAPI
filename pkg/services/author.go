package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/errors"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"fmt"
)

type AuthorServices interface {
	GetAllAuthors() (*dto.ListAuthor, *errors.AppError)
	GetAuthor(id int) (*dto.GetAuthorRes, *errors.AppError)
	CreateAuthor(*dto.CreateAuthorReq) *errors.AppError
	GetAuthorByBook(*dto.GetAuthorByBookReq) (*dto.GetAuthorByBookRes, *errors.AppError)
}

type DefaultAuthor struct {
	repo repository.AuthorRepository
}

func NewAuthor(repo repository.AuthorRepository) AuthorServices {
	return DefaultAuthor{
		repo: repo,
	}
}

func (d DefaultAuthor) GetAllAuthors() (*dto.ListAuthor, *errors.AppError) {
	authors, err := d.repo.List()
	if err != nil {
		return nil, err
	}
	for _, author := range authors {
		fmt.Println(author)
	}
	return &dto.ListAuthor{
		Authors: authors,
	}, nil
}

func (d DefaultAuthor) GetAuthor(id int) (*dto.GetAuthorRes, *errors.AppError) {
	author, err := d.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return &dto.GetAuthorRes{Author: author}, nil

}

func (d DefaultAuthor) CreateAuthor(author *dto.CreateAuthorReq) *errors.AppError {
	result := models.Author{
		Name: author.Name,
	}

	err := d.repo.Create(&result)

	if err != nil {
		panic(err.Error())
	}
	return nil
}

func (d DefaultAuthor) GetAuthorByBook(book *dto.GetAuthorByBookReq) (*dto.GetAuthorByBookRes, *errors.AppError) {
	author, err := d.repo.Get(book.Book.AuthorId)

	if err != nil {
		return nil, err
	}

	return &dto.GetAuthorByBookRes{Author: author}, nil
}
