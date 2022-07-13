package services

import (
	"book-author/pkg/dto"
	"book-author/pkg/models"
	"book-author/pkg/repository"
	"log"
)

type CategoryServices interface {
	CreateCategory(*dto.CreateCateReq) error
	GetCateById(id int) (*dto.GetCateRes, error)
	// GetCateByName(name string) (*dto.GetCateRes, error)
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

// func (d DefaultCategory) GetCateByName(name string) (*dto.GetCateRes, error) {
// 	cate, err := d.repo.GetByName(name)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return &dto.GetCateRes{
// 		Category: cate,
// 	}, nil
// }
