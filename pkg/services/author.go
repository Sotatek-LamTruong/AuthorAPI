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
	GetAuthorsByBook() (*dto.GetAuthorsByBook, error)
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
		fmt.Println(err)
	}
	for _, author := range authors {
		books, err := d.repo.GetBookByAuthor(author.IdAuthor)
		if err != nil {
			fmt.Println(err)
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
		fmt.Println(err)
	}
	books, err := d.repo.GetBookByAuthor(id)
	if err != nil {
		fmt.Println(err)
	}

	return &dto.GetAuthorRes{
		AuthorId:   author.IdAuthor,
		AuthorName: author.Name,
		Books:      books,
	}, nil

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

func (d DefaultAuthor) GetAuthorsByBook() (*dto.GetAuthorsByBook, error) {
	books, err := d.repo.GetAllBooks()
	// book := dto.GetAuthorByBook{}
	var list []dto.GetAuthorByBook
	for _, b := range books {
		book := dto.GetAuthorByBook{}
		authors, err := d.repo.GetByBook(b.IdBook)
		if err != nil {
			fmt.Println(err)
		}
		book.Authors = authors
		book.BookId = b.IdBook
		book.BookName = b.BookName
		book.CategoryId = b.CategoryId

		list = append(list, book)
	}
	if err != nil {
		fmt.Println(err)
	}

	return &dto.GetAuthorsByBook{
		Authors: list,
	}, nil
}
