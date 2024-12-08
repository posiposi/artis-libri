package repository

import (
	"fmt"

	"github.com/posiposi/project/backend/model"
	"gorm.io/gorm"
)

type IBookRepository interface {
	GetAllBooks(books *[]model.Book) error
	GetBookByBookId(book *model.Book, bookId string) error
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book, bookId string) error
	DeleteBook(bookId string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) GetAllBooks(books *[]model.Book) error {
	if err := br.db.Select("books.*, readings.progress_page, readings.review").
		Order("created_at").
		Joins("left join readings on readings.book_id = books.id").
		Find(books).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) GetBookByBookId(book *model.Book, bookId string) error {
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

func (br *bookRepository) UpdateBook(book *model.Book, bookId string) error {
	result := br.db.Model(book).Where("id = ?", bookId).Updates(book)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object is not exist")
	}
	return nil
}

func (br *bookRepository) DeleteBook(bookId string) error {
	result := br.db.Where("id = ?", bookId).Delete(&model.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object is not exist")
	}
	return nil
}
