package repository

import (
	"book-author/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetByAuthor(authId int) ([]models.Book, error)
	GetByCate(cateID int) ([]models.Book, error)
	GetByName(name string) (*models.Book, error)
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
	query := fmt.Sprintf("INSERT INTO `author_book_db`.`book`(`bookname`,`authorId`,`categoryId`) VALUES (%v,%d,%d)", book.Bookname, book.Author.IdAuthor, book.Category.CategoryId)

	result, err := b.db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}

	newId, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The last inserted row id: %d", newId)

	return nil

}

func (b DefaulBookRepository) GetByAuthor(authId int) ([]models.Book, error) {
	var books []models.Book
	book := models.Book{}
	cate := models.Category{}
	author := models.Author{}
	query := fmt.Sprintf("SELECT book.idbook, book.bookname, b.idCategory,b.nameCategory, c.idAuthor,c.Name FROM author_book_db.book as book join author_book_db.category as b on book.categoryId = b.idCategory join author_book_db.author as c on book.authorId = c.idAuthor where idAuthor = %d", authId)

	result, err := b.db.Query(query)

	for result.Next() {
		err := result.Scan(&book.Idbook, &book.Bookname, &cate.CategoryId, &cate.CategoryName, &author.IdAuthor, &author.Name)

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

func (b DefaulBookRepository) GetByCate(cateId int) ([]models.Book, error) {
	var books []models.Book
	book := models.Book{}
	cate := models.Category{}
	author := models.Author{}
	query := fmt.Sprintf("SELECT book.idbook, book.bookname, b.idCategory,b.nameCategory, c.idAuthor,c.Name FROM author_book_db.book as book join author_book_db.category as b on book.categoryId = b.idCategory join author_book_db.author as c on book.authorId = c.idAuthor where categoryId = %d", cateId)

	result, err := b.db.Query(query)

	for result.Next() {
		err := result.Scan(&book.Idbook, &book.Bookname, &cate.CategoryId, &cate.CategoryName, &author.IdAuthor, &author.Name)
		book.Category = cate
		book.Author = author
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

func (b DefaulBookRepository) GetByName(name string) (*models.Book, error) {
	book := models.Book{}
	query, err := b.db.Prepare("SELECT book.idbook, book.bookname, b.nameCategory, c.Name, book.authorId, book.categoryId FROM author_book_db.book as book join author_book_db.category as b on book.categoryId = b.idCategory join author_book_db.author as c on book.authorId = c.idAuthor where bookname like ?")
	if err != nil {
		log.Fatal(err)
	}
	result := query.QueryRow("%" + name + "%")

	err = result.Scan(&book.Idbook, &book.Bookname, &book.Category.CategoryName, &book.Author.Name, &book.Author.IdAuthor, &book.Category.CategoryId)
	if err != nil {
		fmt.Println(err)
	}
	return &book, nil
}
