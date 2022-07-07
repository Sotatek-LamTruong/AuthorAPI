package repository

import (
	"book-author/pkg/errors"
	"book-author/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type AuthorRepository interface {
	List() ([]*models.Author, *errors.AppError)
	Create(author *models.Author) *errors.AppError
	Get(id int) (*models.Author, *errors.AppError)
}

type DefaulAuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) DefaulAuthorRepository {
	return DefaulAuthorRepository{
		db: db,
	}
}

func (r DefaulAuthorRepository) List() ([]*models.Author, *errors.AppError) {
	var authors []*models.Author
	result, err := r.db.Query("SELECT * FROM author_book_db.author")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var author *models.Author
		err = result.Scan(&author.IdAuthor, &author.Name)
		if err != nil {
			panic(err.Error())
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r DefaulAuthorRepository) Create(author *models.Author) *errors.AppError {
	query := fmt.Sprintf("INSERT INTO author_book_db.author (Name) VALUES (%v)", author.Name)

	result, err := r.db.Exec(query)

	if err != nil {
		panic(err.Error())
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d", lastID)
	return nil
}

func (r DefaulAuthorRepository) Get(id int) (*models.Author, *errors.AppError) {
	var author *models.Author
	query := fmt.Sprintf("ISELECT * FROM author_book_db.author where idAuthor = %d", id)
	result, err := r.db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	err = result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		log.Fatal(err)
	}
	return author, nil
}
