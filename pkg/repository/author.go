package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
)

type AuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
	GetAllBooks() ([]models.Book, error)
	GetBookByAuthor(authId int) ([]models.Book, error)
	GetAuthorByBook(bId int) ([]models.Author, error)
	GetByBook(bookID int) ([]models.Author, error)
	Create(author *models.Author) error
	Get(id int) (*models.Author, error)
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
	fmt.Println(result)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		author := models.Author{}
		err := result.Scan(&author.IdAuthor, &author.Name)

		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r DefaulAuthorRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	result, err := r.db.Query("SELECT * FROM author_book_db.book")

	if err != nil {
		return nil, err
	}

	for result.Next() {
		book := models.Book{}
		err := result.Scan(&book.IdBook, &book.BookName, &book.CategoryId)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}

func (b DefaulAuthorRepository) GetAuthorByBook(bId int) ([]models.Author, error) {
	var authors []models.Author
	query := fmt.Sprintf("SELECT bookid,bookname FROM author_book_db.author as a join author_book as b on a.idAuthor = b.authorid join book as c on b.bookid = c.idbook where bookid = %d", bId)

	result, err := b.db.Query(query)

	for result.Next() {
		// cate := models.Category{}
		author := models.Author{}
		err := result.Scan(&author.IdAuthor, &author.Name)

		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err != nil {
		return nil, err
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
			return nil, err
		}

		books = append(books, book)
	}

	if err != nil {
		return nil, err
	}

	return books, nil

}

func (r DefaulAuthorRepository) Create(author *models.Author) error {
	query := fmt.Sprintf("INSERT INTO author_book_db.author (Name) VALUES ('%v')", author.Name)
	fmt.Println(author.Name)
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

func (r DefaulAuthorRepository) GetByBook(id int) ([]models.Author, error) {
	var authors []models.Author
	query := fmt.Sprintf("SELECT authorid,Name FROM author_book_db.author as a join author_book as b on a.idAuthor = b.authorid join book as c on b.bookid = c.idbook where bookid = %d", id)
	result, err := r.db.Query(query)
	// if err != nil {
	// 	fmt.Println("Fail")
	// }

	for result.Next() {
		author := models.Author{}
		err := result.Scan(&author.IdAuthor, &author.Name)
		if err != nil {
			fmt.Println("Scan fail")
		}
		authors = append(authors, author)
	}
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r DefaulAuthorRepository) Get(id int) (*models.Author, error) {
	var author = new(models.Author)
	query := fmt.Sprintf("SELECT * FROM author_book_db.author where idAuthor = %d;", id)
	result := r.db.QueryRow(query)

	err := result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		return nil, err
	}
	return author, nil
}
