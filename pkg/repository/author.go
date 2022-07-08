package repository

import (
	"book-author/pkg/errors"
	"book-author/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type AuthorRepository interface {
	List() ([]models.Author, *errors.AppError)
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

func (r DefaulAuthorRepository) List() ([]models.Author, *errors.AppError) {
	var authors []models.Author
	result, err := r.db.Query("SELECT idAuthor,Name FROM author_book_db.author;")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()
	for result.Next() {
		var author models.Author
		err := result.Scan(&author.IdAuthor, &author.Name)
		if err != nil {
			panic(err.Error())
		}
		authors = append(authors, author)
		fmt.Println(authors)
	}
	return authors, nil
}

func (r DefaulAuthorRepository) Create(author *models.Author) *errors.AppError {
	query := fmt.Sprintf("INSERT INTO author_book_db.author (Name) VALUES ('%v')", author.Name)
	fmt.Println(author.Name)
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

func (r DefaulAuthorRepository) Get(id int) (*models.Author, *errors.AppError) {
	var author = new(models.Author)
	query := fmt.Sprintf("SELECT * FROM author_book_db.author where idAuthor = %d", id)
	result := r.db.QueryRow(query)

	err := result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		log.Fatal(err)
	}
	return author, nil
}

func (r DefaulAuthorRepository) getByBook(id int) (*models.Author, *errors.AppError) {
	var author = new(models.Author)
	query := fmt.Sprintf("SELECT idAuthor,Name from author_book_db.book as a join author_book_db.author as b on a.authorId = b.idAuthor where idbook = %d", id)
	result := r.db.QueryRow(query)

	err := result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		log.Fatal(err)
	}
	return author, nil
}
