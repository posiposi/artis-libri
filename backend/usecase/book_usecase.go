package usecase

import (
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/repository"
)

type IBookUsecase interface {
	GetAllBooks() ([]model.BookResponse, error)
	GetBookByBookId(bookId uint) (model.BookResponse, error)
	CreateBook(book model.Book) (model.BookResponse, error)
	UpdateBook(book model.Book, bookId uint) (model.BookResponse, error)
	DeleteBook(bookId uint) error
}

type bookUsecase struct {
	br repository.IBookRepository
}

func NewBookUsecase(br repository.IBookRepository) IBookUsecase {
	return &bookUsecase{br}
}

func (bu *bookUsecase) GetAllBooks() ([]model.BookResponse, error) {
	books := []model.Book{}
	if err := bu.br.GetAllBooks(&books); err != nil {
		return nil, err
	}
	bookRess := []model.BookResponse{}
	for _, v := range books {
		t := model.BookResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		bookRess = append(bookRess, t)
	}
	return bookRess, nil
}

func (bu *bookUsecase) GetBookByBookId(bookId uint) (model.BookResponse, error) {
	book := model.Book{}
	if err := bu.br.GetBookByBookId(&book, bookId); err != nil {
		return model.BookResponse{}, err
	}
	bookRes := model.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return bookRes, nil
}

func (bu *bookUsecase) CreateBook(book model.Book) (model.BookResponse, error) {
	if err := bu.br.CreateBook(&book); err != nil {
		return model.BookResponse{}, err
	}
	bookRes := model.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return bookRes, nil
}

func (bu *bookUsecase) UpdateBook(book model.Book, bookId uint) (model.BookResponse, error) {
	if err := bu.br.UpdateBook(&book, bookId); err != nil {
		return model.BookResponse{}, err
	}
	bookRes := model.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return bookRes, nil
}

func (bu *bookUsecase) DeleteBook(bookId uint) error {
	if err := bu.br.DeleteBook(bookId); err != nil {
		return err
	}
	return nil
}
