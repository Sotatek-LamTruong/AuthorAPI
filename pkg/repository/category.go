package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type CategoryRepository interface {
	Create(cate *models.Category) error
	GetByBook(id int) (*models.Category, error)
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

func (r DefaulCateRepository) GetByBook(bookID int) (*models.Category, error) {
	cate := models.Category{}
	query := fmt.Sprintf("SELECT idCategory,nameCategory from author_book_db.book as a join author_book_db.category as b on a.categoryId = b.idCategory where idbook = %d", bookID)
	result := r.db.QueryRow(query)

	err := result.Scan(&cate.CategoryId, &cate.CategoryName)
	if err != nil {
		return nil, err
	}
	return &cate, nil
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
