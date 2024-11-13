package repository

import (
	"fmt"

	"github.com/posiposi/project/backend/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBookRepository interface {
	GetAllBooks(books *[]model.Book) error
	GetBookByBookId(book *model.Book, bookId uint) error
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book, bookId uint) error
	DeleteBook(bookId uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) GetAllBooks(books *[]model.Book) error {
	if err := br.db.Order("created_at").Find(books).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) GetBookByBookId(book *model.Book, bookId uint) error {
	if err := br.db.First(book, bookId).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) CreateBook(book *model.Book) error {
	if err := br.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) UpdateBook(book *model.Book, bookId uint) error {
	result := br.db.Model(book).Clauses(clause.Returning{}).Where("book_id = ?", bookId).Update("title", book.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object is not exist")
	}
	return nil
}

func (br *bookRepository) DeleteBook(bookId uint) error {
	result := br.db.Where("book_id = ?", bookId).Delete(&model.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object is not exist")
	}
	return nil
}
