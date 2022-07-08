package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"log"
)

type CategoryServices interface {
	CreateCategory(*dto.CreateCateReq) error
	GetCateByBook(bookID int) (*dto.GetCateByBookRes, error)
	GetCateByName(name string) (*dto.GetCateByNameRes, error)
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

func (d DefaultCategory) GetCateByBook(bookID int) (*dto.GetCateByBookRes, error) {
	cate, err := d.repo.GetByBook(bookID)

	if err != nil {
		log.Fatalln(err)
	}

	return &dto.GetCateByBookRes{
		Category: cate,
	}, nil
}

func (d DefaultCategory) GetCateByName(name string) (*dto.GetCateByNameRes, error) {
	cate, err := d.repo.GetByName(name)

	if err != nil {
		log.Fatalln(err)
	}

	return &dto.GetCateByNameRes{
		Category: cate,
	}, nil
}
