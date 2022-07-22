package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
)

type AuthorRepository interface {
	Create(author *models.Author) (int64, error)
	Get(id int) (*models.Author, error)
	GetBooks(author *models.Author) ([]models.Book, error)
	GetAllAuthors() ([]models.Author, error)
	// GetAllBooks() ([]models.Book, error)
	// GetBookByAuthor(authId int) ([]models.Book, error)
	GetBook(id int) (*models.Book, error)
	GetAuthors(bo *models.Book) ([]models.Author, error)
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

	result, err := r.db.Query("SELECT * FROM author")
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

func (r DefaulAuthorRepository) GetBook(id int) (*models.Book, error) {
	book := models.Book{}
	query := fmt.Sprintf("SELECT * FROM newdb.book where bookid = %d", id)

	result := r.db.QueryRow(query)

	err := result.Scan(&book.IdBook, &book.BookName)

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b DefaulAuthorRepository) GetAuthors(bo *models.Book) ([]models.Author, error) {
	var authors []models.Author
	query := fmt.Sprintf("select a.authorid, a.name from author as a join author_book as b on a.authorid = b.authorid join book as c on c.bookid = b.bookid where c.bookid = %d", bo.IdBook)

	result, err := b.db.Query(query)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		author := models.Author{}
		err := result.Scan(&author.IdAuthor, &author.Name)

		fmt.Println(authors)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil

}

// func (b DefaulAuthorRepository) GetBookByAuthor(authId int) ([]models.Book, error) {
// 	var books []models.Book
// 	query := fmt.Sprintf("SELECT bookid,bookname,categoryId FROM author_book_db.author as a join author_book as b on a.idAuthor = b.authorid join book as c on b.bookid = c.idbook where idAuthor = %d", authId)

// 	result, err := b.db.Query(query)

// 	for result.Next() {
// 		// cate := models.Category{}
// 		book := models.Book{}
// 		err := result.Scan(&book.IdBook, &book.BookName, &book.CategoryId)

// 		if err != nil {
// 			return nil, err
// 		}

// 		books = append(books, book)
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	return books, nil

// }

func (r DefaulAuthorRepository) Create(author *models.Author) (int64, error) {
	query := fmt.Sprintf("INSERT INTO author (Name) VALUES ('%v')", author.Name)
	fmt.Println(author.Name)
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

func (r DefaulAuthorRepository) GetBooks(author *models.Author) ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf("select c.bookid,c.name from author as a join author_book as b on a.authorid = b.authorid join book as c on c.bookid = b.bookid where a.authorid = %d", author.IdAuthor)
	result, err := r.db.Query(query)
	for result.Next() {
		book := models.Book{}
		err := result.Scan(&book.IdBook, &book.BookName)
		if err != nil {
			fmt.Println("Scan fail")
		}
		books = append(books, book)
	}
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r DefaulAuthorRepository) Get(id int) (*models.Author, error) {
	var author = new(models.Author)
	fmt.Println(id)
	query := fmt.Sprintf("SELECT * FROM author where authorid = %d;", id)
	result := r.db.QueryRow(query)

	err := result.Scan(&author.IdAuthor, &author.Name)
	if err != nil {
		return nil, err
	}
	return author, nil
}
