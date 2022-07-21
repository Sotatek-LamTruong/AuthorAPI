package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
)

type BookRepository interface {
	Create(book *models.Book) (int64, error)
	Get(id int) (*models.Book, error)
	AddCate(book *models.Book, cateId int) (int64, error)
	AddAuthor(book *models.Book, authId int) (int64, error)
	DeleteAuthor(book *models.Book, authId int) (int64, error)
	UpdateAuthor(authId int, name string) (int64, error)
	GetByAuthor(authId int) (*models.Author, []models.Book, error)
	GetByCate(cateId int) (*models.Category, []models.Book, error)
	GetCategories(bookid int) (*models.Book, []models.Category, error)
	GetAuthors(bookid int) (*models.Book, []models.Author, error)
}

type DefaulBookRepository struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) DefaulBookRepository {
	return DefaulBookRepository{
		db: db,
	}
}

func (b DefaulBookRepository) Create(book *models.Book) (int64, error) {
	fmt.Println(book)
	query := fmt.Sprintf("INSERT INTO book (name) VALUES ('%v');", book.BookName)

	result, err := b.db.Exec(query)

	if err != nil {
		return 0, err
	}

	newId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	fmt.Printf("The last inserted row id: %d", newId)

	return newId, nil

}

func (b DefaulBookRepository) Get(id int) (*models.Book, error) {
	book := models.Book{}
	query := fmt.Sprintf("SELECT * FROM newdb.book where bookid = %d", id)

	result := b.db.QueryRow(query)

	err := result.Scan(&book.IdBook, &book.BookName)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b DefaulBookRepository) AddCate(book *models.Book, cateId int) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `category_book`(`categoryid`, `bookid`) VALUES (%d,%d);", cateId, book.IdBook)

	result, err := b.db.Exec(query)

	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (b DefaulBookRepository) AddAuthor(book *models.Book, authId int) (int64, error) {
	query := fmt.Sprintf("INSERT INTO `author_book`(`authorid`, `bookid`) VALUES (%d,%d);", authId, book.IdBook)

	result, err := b.db.Exec(query)

	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (b DefaulBookRepository) DeleteAuthor(book *models.Book, authId int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM `newdb`.`author_book` WHERE bookid = %d and authorid = %d", book.IdBook, authId)

	result, err := b.db.Exec(query)

	if err != nil {
		return 0, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (b DefaulBookRepository) UpdateAuthor(authId int, name string) (int64, error) {
	query := fmt.Sprintf("UPDATE author SET name = '%v' WHERE authorid = %d", name, authId)

	result, err := b.db.Exec(query)
	if err != nil {
		return 0, err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (b DefaulBookRepository) GetByAuthor(authId int) (*models.Author, []models.Book, error) {
	var books []models.Book
	author := models.Author{}
	query := fmt.Sprintf("select a.bookid,a.name,c.authorid,c.name from book as a join author_book as b on a.bookid = b.bookid join author as c on c.authorid = b.authorid where c.authorid = %d", authId)

	result, err := b.db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	for result.Next() {
		book := models.Book{}
		err := result.Scan(&book.IdBook, &book.BookName, &author.IdAuthor, &author.Name)

		if err != nil {
			return nil, nil, err
		}

		books = append(books, book)

	}

	if err != nil {
		fmt.Println(err)
	}

	return &author, books, nil

}

func (b DefaulBookRepository) GetByCate(cateId int) (*models.Category, []models.Book, error) {
	var books []models.Book
	cate := models.Category{}
	query := fmt.Sprintf("select a.bookid,a.name,c.categoryid,c.name from book as a join category_book as b on a.bookid = b.bookid join category as c on c.categoryid = b.categoryid where c.categoryid = %d", cateId)

	result, err := b.db.Query(query)

	for result.Next() {
		book := models.Book{}
		err := result.Scan(&book.IdBook, &book.BookName, &cate.CategoryId, &cate.CategoryName)

		if err != nil {
			return nil, nil, err
		}

		books = append(books, book)

	}

	if err != nil {
		return nil, nil, err
	}

	return &cate, books, nil

}

func (b DefaulBookRepository) GetCategories(bookid int) (*models.Book, []models.Category, error) {
	var categories []models.Category
	book, err := b.Get(bookid)
	if err != nil {
		return nil, nil, err
	}
	query := fmt.Sprintf("select c.categoryid,c.name from book as a join category_book as b on a.bookid = b.bookid join category as c on c.categoryid = b.categoryid where a.bookid = %d", bookid)
	result, err := b.db.Query(query)
	if err != nil {
		return nil, nil, err
	}

	for result.Next() {
		cate := models.Category{}

		err := result.Scan(&cate.CategoryId, &cate.CategoryName)
		if err != nil {
			return nil, nil, err
		}

		categories = append(categories, cate)
	}

	return book, categories, nil
}

func (b DefaulBookRepository) GetAuthors(bookid int) (*models.Book, []models.Author, error) {
	var authors []models.Author
	book, err := b.Get(bookid)
	if err != nil {
		return nil, nil, err
	}
	query := fmt.Sprintf("SELECT c.authorid, c.name FROM newdb.book as a join author_book as b on a.bookid = b.bookid join author as c on c.authorid = b.authorid where a.bookid = %d", bookid)
	result, err := b.db.Query(query)
	if err != nil {
		return nil, nil, err
	}

	for result.Next() {
		author := models.Author{}

		err := result.Scan(&author.IdAuthor, &author.Name)
		if err != nil {
			return nil, nil, err
		}

		authors = append(authors, author)
	}

	return book, authors, nil
}
