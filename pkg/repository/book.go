package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetByAuthor(authId int) ([]models.Book, error)
	GetByCate(cateID int) ([]models.Book, error)
	GetByName(name string) (*models.Book, error)
	UpdateAuthor(aId int, name string) error
	GetAuthorByBook(BookId int, AuthorId int) (*models.Author, error)
}

type DefaulBookRepository struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) DefaulBookRepository {
	return DefaulBookRepository{
		db: db,
	}
}

func (b DefaulBookRepository) Create(book *models.Book) error {
	fmt.Println(book)
	query := fmt.Sprintf("INSERT INTO book (bookname) VALUES ('%v');", book.BookName)

	result, err := b.db.Exec(query)

	if err != nil {
		return err
	}

	newId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Printf("The last inserted row id: %d", newId)

	return nil

}

func (b DefaulBookRepository) UpdateAuthor(aId int, name string) error {
	query := fmt.Sprintf("UPDATE author SET Name = '%v' WHERE idAuthor = %d", name, aId)

	_, err := b.db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Update success")
	return nil
}

func (b DefaulBookRepository) Get(aId int, name string) error {
	query := fmt.Sprintf("UPDATE author SET Name = '%v' WHERE idAuthor = %d", name, aId)

	_, err := b.db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Update success")
	return nil
}

func (b DefaulBookRepository) GetAuthorByBook(BookId int, AuthorId int) (*models.Author, error) {
	author := models.Author{}

	query := fmt.Sprintf("SELECT idAuthor,Name,bookid,bookname,categoryId FROM author as a join author_book as b on a.idAuthor = b.authorid join book as c on c.idbook = b.bookid WHERE idAuthor = %d AND bookid = %d", AuthorId, BookId)

	result := b.db.QueryRow(query)

	result.Scan(&author.IdAuthor, &author.Name)
	return &author, nil
}

func (b DefaulBookRepository) GetByAuthor(authId int) ([]models.Book, error) {
	var books []models.Book
	book := models.Book{}
	cate := models.Category{}
	author := models.Author{}
	query := fmt.Sprintf("SELECT book.idbook, book.bookname, b.idCategory,b.nameCategory, c.idAuthor,c.Name FROM author_book_db.book as book join author_book_db.category as b on book.categoryId = b.idCategory join author_book_db.author as c on book.authorId = c.idAuthor where idAuthor = %d", authId)

	result, err := b.db.Query(query)

	for result.Next() {
		err := result.Scan(&book.IdBook, &book.BookName, &cate.CategoryId, &cate.CategoryName, &author.IdAuthor, &author.Name)

		if err != nil {
			return nil, err
		}

		books = append(books, book)

	}

	if err != nil {
		fmt.Println(err)
	}

	return books, nil

}

func (b DefaulBookRepository) GetByCate(cateId int) ([]models.Book, error) {
	var books []models.Book
	book := models.Book{}
	cate := models.Category{}
	author := models.Author{}
	query := fmt.Sprintf("SELECT book.idbook, book.bookname, b.idCategory,b.nameCategory, c.idAuthor, c.Name FROM author_book_db.book as book join author_book_db.category as b on book.categoryId = b.idCategory join author_book_db.author as c on book.authorId = c.idAuthor where categoryId = %d", cateId)

	result, err := b.db.Query(query)

	for result.Next() {
		err := result.Scan(&book.IdBook, &book.BookName, &cate.CategoryId, &cate.CategoryName, &author.IdAuthor, &author.Name)

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

func (b DefaulBookRepository) GetByName(name string) (*models.Book, error) {
	book := models.Book{}
	query, err := b.db.Prepare("SELECT book.idbook, book.bookname, b.nameCategory, c.Name, book.authorId, book.categoryId FROM author_book_db.book as book join author_book_db.category as b on book.categoryId = b.idCategory join author_book_db.author as c on book.authorId = c.idAuthor where bookname like ?")
	if err != nil {
		return nil, err
	}
	result := query.QueryRow("%" + name + "%")

	err = result.Scan(&book.IdBook, &book.BookName)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
