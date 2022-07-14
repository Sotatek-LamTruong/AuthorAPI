package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type AuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
	GetBookByAuthor(authId int) ([]models.Book, error)
	Create(author *models.Author) error
	Get(id int) (*models.Author, error)
	GetByBook(bookID int) (*models.Author, error)
}

type DefaulAuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepo(db *sql.DB) DefaulAuthorRepository {
	return DefaulAuthorRepository{
		db: db,
	}
}

func (r DefaulAuthorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author

	result, err := r.db.Query("SELECT * FROM author_book_db.author")

	if err != nil {
		fmt.Println(err)
	}

	for result.Next() {
		author := models.Author{}
		err := result.Scan(&author.IdAuthor, &author.Name)

		if err != nil {
			log.Fatal(err)
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (b DefaulAuthorRepository) GetBookByAuthor(authId int) ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf("SELECT bookid,bookname,categoryId FROM author_book_db.author as a join author_book as b on a.idAuthor = b.authorid join book as c on b.bookid = c.idbook where idAuthor = %d", authId)

	result, err := b.db.Query(query)

	for result.Next() {
		// cate := models.Category{}
		book := models.Book{}
		err := result.Scan(&book.IdBook, &book.BookName, &book.CategoryId)

		if err != nil {
			fmt.Println(err)
		}

		books = append(books, book)
	}

	if err != nil {
		fmt.Println(err)
	}

	return books, nil

}

func (r DefaulAuthorRepository) List() ([]models.Author, []models.Book, error) {
	authors, err := r.GetAllAuthors()
	if err != nil {
		fmt.Println(err)
	}
	var books []models.Book
	result, err := r.db.Query("SELECT idAuthor,Name,bookid,bookname,categoryId FROM author as a join author_book as b on a.idAuthor = b.authorid join book as c on c.idbook = b.bookid join category as d on d.idCategory = c.categoryId")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()
	for result.Next() {
		author := models.Author{}
		book := models.Book{}
		err := result.Scan(&author.IdAuthor, &author.Name, &book.IdBook, &book.BookName, &book.CategoryId)
		if err != nil {
			panic(err.Error())
		}
		authors = append(authors, author)
		books = append(books, book)
		fmt.Println(authors)
	}
	return authors, books, nil
}

func (r DefaulAuthorRepository) Create(author *models.Author) error {
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

func (r DefaulAuthorRepository) Get(id int) (*models.Author, error) {
	var author = new(models.Author)
	query := fmt.Sprintf("SELECT * FROM author_book_db.author where idAuthor = %d", id)
	result := r.db.QueryRow(query)

	err := result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		log.Fatal(err)
	}
	return author, nil
}

func (r DefaulAuthorRepository) GetByBook(id int) (*models.Author, error) {
	var author = new(models.Author)
	query := fmt.Sprintf("SELECT idAuthor,Name from author_book_db.book as a join author_book_db.author as b on a.authorId = b.idAuthor where idbook = %d", id)
	result := r.db.QueryRow(query)

	err := result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		log.Fatal(err)
	}
	return author, nil
}
