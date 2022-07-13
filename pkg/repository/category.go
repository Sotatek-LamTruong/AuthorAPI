package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type CategoryRepository interface {
	Create(cate *models.Category) error
	GetById(id int) (*models.Category, []models.Book, error)
	GetByName(name string) (*models.Category, error)
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
		log.Fatal(err)
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		fmt.Println(book)
		books = append(books, book)
	}
	return &cate, books, nil
}
func (r DefaulCateRepository) GetByName(name string) (*models.Category, error) {
	fmt.Println(name)
	var cate = new(models.Category)
	query, err := r.db.Prepare("SELECT * FROM author_book_db.category where nameCategory like ? ")
	if err != nil {
		log.Fatal(err)
	}
	result := query.QueryRow("%" + name + "%")
	err = result.Scan(&cate.CategoryId, &cate.CategoryName)
	if err != nil {
		fmt.Println(err)
	}
	return cate, nil
}
