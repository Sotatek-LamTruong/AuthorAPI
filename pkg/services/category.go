package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"fmt"
)

type CategoryServices interface {
	CreateCategory(*dto.CreateCateReq) (int64, error)
	GetCate(id int) (*dto.GetCateRes, error)
	GetCateByBook(id int) (*dto.GetCateRes, error)
	DeleteCategory(id int) (int64, error)
}

type DefaultCategory struct {
	repo repository.CategoryRepository
}

func NewCategory(repo repository.CategoryRepository) CategoryServices {
	return DefaultCategory{
		repo: repo,
	}
}

func (d DefaultCategory) CreateCategory(cate *dto.CreateCateReq) (int64, error) {
	result := models.Category{
		CategoryName: cate.CategoryName,
	}
	row, err := d.repo.Create(&result)
	if err != nil {
		return 0, err
	}

	return row, nil
}

func (d DefaultCategory) GetCate(Id int) (*dto.GetCateRes, error) {
	cate, err := d.repo.Get(Id)
	if err != nil {
		return nil, err
	}
	books, err := d.repo.GetBooks(cate)
	fmt.Println(books)
	if err != nil {
		return nil, err
	}

	return &dto.GetCateRes{
		CategoryId:   cate.CategoryId,
		CategoryName: cate.CategoryName,
		Books:        books,
	}, nil
}

func (d DefaultCategory) GetCateByBook(id int) (*dto.GetCateRes, error) {
	fmt.Println(id)
	cate, err := d.repo.GetByBook(id)

	if err != nil {
		return nil, err
	}

	return &dto.GetCateRes{
		CategoryId:   cate.CategoryId,
		CategoryName: cate.CategoryName,
	}, nil
}

func (d DefaultCategory) DeleteCategory(id int) (int64, error) {
	result, err := d.repo.Delete(id)

	if err != nil {
		return 0, err
	}

	return result, nil
}
