package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"fmt"
)

type BookServices interface {
	CreateBook(book *dto.AddBookReq) (int64, error)
	AddCategory(bookid int, req *dto.CateReq) (*dto.AuthorRes, error)
	AddAuthors(bookid int, req *dto.AuthorReq) (*dto.AuthorRes, error)
	EditAuthor(bookid int, req *dto.AuthorReq) (*dto.AuthorRes, error)
	DeleteAuthor(bookid int, req *dto.AuthorReq) (*dto.AuthorRes, error)
	GetBooksByAuthor(req *dto.AuthorReq) (*dto.GetByAuthor, error)
	GetBooksByCate(req *dto.CateReq) (*dto.GetByCate, error)
}

type DefaultBook struct {
	repo repository.BookRepository
}

func NewBook(repo repository.BookRepository) BookServices {
	return DefaultBook{
		repo: repo,
	}
}

func (b DefaultBook) CreateBook(book *dto.AddBookReq) (int64, error) {
	result := models.Book{
		BookName: book.BookName,
	}
	fmt.Println(result)

	row, err := b.repo.Create(&result)

	if err != nil {
		return 0, err
	}

	return row, nil
}

func (b DefaultBook) AddCategory(bookid int, req *dto.CateReq) (*dto.AuthorRes, error) {
	book, err := b.repo.Get(bookid)
	if err != nil {
		return nil, err
	}
	row, err := b.repo.AddCate(book, req.CategoryID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthorRes{
		Row:      row,
		BookID:   book.IdBook,
		BookName: book.BookName,
	}, nil

}

func (b DefaultBook) AddAuthors(bookid int, req *dto.AuthorReq) (*dto.AuthorRes, error) {
	book, err := b.repo.Get(bookid)
	if err != nil {
		return nil, err
	}
	row, err := b.repo.AddAuthor(book, req.AuthorID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthorRes{
		Row:      row,
		BookID:   book.IdBook,
		BookName: book.BookName,
	}, nil

}
func GetAuthor(authors []models.Author, authid int) *models.Author {
	for _, author := range authors {
		if author.IdAuthor == authid {
			return &author
		}
	}
	return nil
}

func (b DefaultBook) EditAuthor(bookid int, req *dto.AuthorReq) (*dto.AuthorRes, error) {

	book, authors, err := b.repo.GetAuthors(bookid)
	if err != nil {
		return nil, err
	}
	author := GetAuthor(authors, req.AuthorID)
	if author == nil {
		return nil, nil
	}
	row, err := b.repo.UpdateAuthor(req.AuthorID, req.Name)
	if err != nil {
		return nil, err
	}
	return &dto.AuthorRes{
		Row:      row,
		BookID:   book.IdBook,
		BookName: book.BookName,
	}, nil

}

func (b DefaultBook) DeleteAuthor(bookid int, req *dto.AuthorReq) (*dto.AuthorRes, error) {

	book, authors, err := b.repo.GetAuthors(bookid)
	if err != nil {
		return nil, err
	}
	author := GetAuthor(authors, req.AuthorID)
	if author == nil {
		fmt.Println("Id not exist")
		return nil, nil
	}
	row, err := b.repo.DeleteAuthor(book, req.AuthorID)
	if err != nil {
		return nil, err
	}
	return &dto.AuthorRes{
		Row:      row,
		BookID:   book.IdBook,
		BookName: book.BookName,
	}, nil

}

func (b DefaultBook) GetBooksByAuthor(req *dto.AuthorReq) (*dto.GetByAuthor, error) {
	author, books, err := b.repo.GetByAuthor(req.AuthorID)
	if err != nil {
		return nil, err
	}

	return &dto.GetByAuthor{
		AuthorID:   author.IdAuthor,
		AuthorName: author.Name,
		Books:      books,
	}, nil
}

func (b DefaultBook) GetBooksByCate(req *dto.CateReq) (*dto.GetByCate, error) {
	author, books, err := b.repo.GetByCate(req.CategoryID)
	if err != nil {
		return nil, err
	}

	return &dto.GetByCate{
		CateID:   author.CategoryId,
		CateName: author.CategoryName,
		Books:    books,
	}, nil
}
