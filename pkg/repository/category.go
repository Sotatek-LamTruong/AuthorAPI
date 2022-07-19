package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
)

type CategoryRepository interface {
	Create(cate *models.Category) error
	GetById(id int) (*models.Category, []models.Book, error)
	GetByBook(id int) (*models.Category, error)
	Delete(id int) error
}

type DefaulCateRepository struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) DefaulCateRepository {
	return DefaulCateRepository{
		db: db,
	}
}

func (r DefaulCateRepository) Create(cate *models.Category) error {
	query := fmt.Sprintf(" INSERT INTO author_book_db.category (nameCategory) VALUES	('%v')", cate.CategoryName)
	fmt.Println(cate.CategoryName)
	result, err := r.db.Exec(query)

	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Printf("The last inserted row id: %d", lastID)
	return nil

}

func (r DefaulCateRepository) GetById(id int) (*models.Category, []models.Book, error) {
	cate := models.Category{}
	var books []models.Book

	query := fmt.Sprintf("SELECT categoryId, nameCategory, idbook, bookname FROM author_book_db.book as a join author_book_db.category as b on a.categoryId = b.idCategory where categoryId = %d", id)
	result, err := r.db.Query(query)

	if err != nil {
		return nil, nil, err
	}
	for result.Next() {
		book := models.Book{}

		err := result.Scan(&cate.CategoryId, &cate.CategoryName, &book.IdBook, &book.BookName)
		if err != nil {
			return nil, nil, err
		}
		fmt.Println(book)
		books = append(books, book)
	}
	return &cate, books, nil
}
func (r DefaulCateRepository) GetByBook(id int) (*models.Category, error) {
	var cate = models.Category{}
	fmt.Println(id)
	query := fmt.Sprintf("SELECT idCategory,nameCategory from category as a join book as b on a.idCategory = b.idbook where b.idbook = %d", id)
	result := r.db.QueryRow(query)
	err := result.Scan(&cate.CategoryId, &cate.CategoryName)
	fmt.Println(cate)
	if err != nil {
		return nil, err
	}
	return &cate, nil
}

func (r DefaulCateRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM category where idCategory = %d", id)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
