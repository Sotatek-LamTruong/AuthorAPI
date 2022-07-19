package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
)

type CategoryRepository interface {
	Create(cate *models.Category) (int64, error)
	Get(id int) (*models.Category, error)
	GetByBook(id int) (*models.Category, error)
	Delete(id int) (int64, error)
	GetBooks(category *models.Category) ([]models.Book, error)
}

type DefaulCateRepository struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) DefaulCateRepository {
	return DefaulCateRepository{
		db: db,
	}
}

func (r DefaulCateRepository) Create(cate *models.Category) (int64, error) {
	query := fmt.Sprintf(" INSERT INTO category (name) VALUES	('%v')", cate.CategoryName)
	result, err := r.db.Exec(query)

	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	fmt.Printf("The last inserted row id: %d", lastID)
	return lastID, nil

}

func (r DefaulCateRepository) Get(id int) (*models.Category, error) {
	cate := models.Category{}
	query := fmt.Sprintf("SELECT * FROM newdb.category where categoryid = %d;", id)
	result := r.db.QueryRow(query)
	err := result.Scan(&cate.CategoryId, &cate.CategoryName)
	if err != nil {
		return nil, err
	}
	return &cate, nil
}

func (r DefaulCateRepository) GetBooks(category *models.Category) ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf("SELECT c.bookid,c.name FROM `newdb`.`category` as a join category_book as b on a.categoryid = b.categoryid join book as c on b.bookid = c.bookid where a.categoryid = %d", category.CategoryId)
	result, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		book := models.Book{}

		err := result.Scan(&book.IdBook, &book.BookName)
		if err != nil {
			return nil, err
		}
		fmt.Println(book)
		books = append(books, book)
	}
	return books, nil
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

func (r DefaulCateRepository) Delete(id int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM category where categoryid = %d", id)
	result, err := r.db.Exec(query) // check result if id not exist
	if err != nil {
		return 0, err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		fmt.Println("Id not exist")
		return 0, nil
	}
	return row, nil
}
