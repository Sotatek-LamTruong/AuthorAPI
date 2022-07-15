package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"fmt"
	"log"
)

type CategoryServices interface {
	CreateCategory(*dto.CreateCateReq) error
	GetCateById(id int) (*dto.GetCateRes, error)
	GetCateByBook(id int) (*dto.GetCateRes, error)
}

type DefaultCategory struct {
	repo repository.CategoryRepository
}

func NewCategory(repo repository.CategoryRepository) CategoryServices {
	return DefaultCategory{
		repo: repo,
	}
}

func (d DefaultCategory) CreateCategory(cate *dto.CreateCateReq) error {
	result := models.Category{
		CategoryName: cate.CategoryName,
	}
	err := d.repo.Create(&result)
	if err != nil {
		return err
	}

	return nil
}

func (d DefaultCategory) GetCateById(Id int) (*dto.GetCateRes, error) {
	cate, books, err := d.repo.GetById(Id)

	if err != nil {
		log.Fatalln(err)
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
