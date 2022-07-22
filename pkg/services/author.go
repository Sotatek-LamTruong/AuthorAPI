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
	CreateAuthor(dto.CreateAuthorReq) (int64, error)
	GetAuthorsByBook(id int) (*dto.GetAuthorsByBook, error)
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
	authors, err := d.repo.GetAllAuthors()
	auth := dto.AuthorDTO{}
	var list []dto.AuthorDTO
	if err != nil {
		fmt.Println("1")
		return nil, err
	}
	for _, author := range authors {
		books, err := d.repo.GetBooks(&author)
		if err != nil {
			fmt.Println("2")
			return nil, err
		}
		auth.AuthorId = author.IdAuthor
		auth.AuthorName = author.Name
		auth.Books = books
		list = append(list, auth)
	}
	return &dto.ListAuthor{
		Authors: list,
	}, nil
}

func (d DefaultAuthor) GetAuthor(id int) (*dto.GetAuthorRes, error) {
	author, err := d.repo.Get(id)
	if err != nil {
		return nil, err
	}
	books, err := d.repo.GetBooks(author)
	if err != nil {
		return nil, err
	}

	return &dto.GetAuthorRes{
		AuthorId:   author.IdAuthor,
		AuthorName: author.Name,
		Books:      books,
	}, nil

}

func (d DefaultAuthor) CreateAuthor(author dto.CreateAuthorReq) (int64, error) {
	result := models.Author{
		Name: author.Name,
	}

	row, err := d.repo.Create(&result)

	if err != nil {
		return 0, err
	}
	return row, nil
}

func (d DefaultAuthor) GetAuthorsByBook(id int) (*dto.GetAuthorsByBook, error) {
	book, err := d.repo.GetBook(id)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}

	authors, err := d.repo.GetAuthors(book)

	if err != nil {
		fmt.Println("2")
		return nil, err
	}

	return &dto.GetAuthorsByBook{
		BookId:   book.IdBook,
		BookName: book.BookName,
		Authors:  authors,
	}, nil
}
