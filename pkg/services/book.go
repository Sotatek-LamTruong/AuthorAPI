package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"fmt"
)

type BookServices interface {
	CreateBook(book *dto.AddBookReq) error
	GetBookByCate(bookId int) (*dto.GetBookRes, error)
	GetBookByAuthor(authorId int) (*dto.GetBookRes, error)
	GetBookByName(name string) (*dto.GetBookByNameRes, error)
}

type DefaultBook struct {
	repo repository.BookRepository
}

func NewBook(repo repository.BookRepository) BookServices {
	return DefaultBook{
		repo: repo,
	}
}

func (b DefaultBook) CreateBook(book *dto.AddBookReq) error {
	result := models.Book{
		BookName:   book.BookName,
		CategoryId: book.CategoryId,
	}
	fmt.Println(result)

	err := b.repo.Create(&result)

	if err != nil {
		return err
	}

	return nil
}

func (b DefaultBook) GetBookByCate(cateId int) (*dto.GetBookRes, error) {
	books, err := b.repo.GetByCate(cateId)

	if err != nil {
		fmt.Println("Fail")
	}

	return &dto.GetBookRes{
		Books: books,
	}, nil

}

func (b DefaultBook) GetBookByAuthor(authorId int) (*dto.GetBookRes, error) {
	books, err := b.repo.GetByAuthor(authorId)

	if err != nil {
		fmt.Println("Fail")
	}

	return &dto.GetBookRes{
		Books: books,
	}, nil

}

func (b DefaultBook) GetBookByName(name string) (*dto.GetBookByNameRes, error) {
	book, err := b.repo.GetByName(name)

	if err != nil {
		fmt.Println("Fail")
	}

	return &dto.GetBookByNameRes{
		Book: book,
	}, nil
}
