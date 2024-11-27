package usecase

import (
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/repository"
)

type IBookUsecase interface {
	GetAllBooks() ([]model.BookResponse, error)
	GetBookByBookId(bookId string) (model.BookResponse, error)
	CreateBook(book model.Book) (model.BookResponse, error)
	UpdateBook(book model.Book, bookId string) (model.BookResponse, error)
	DeleteBook(bookId string) error
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
	bookRes := []model.BookResponse{}
	for _, v := range books {
		t := model.BookResponse(v)
		bookRes = append(bookRes, t)
	}
	return bookRes, nil
}

func (bu *bookUsecase) GetBookByBookId(bookId string) (model.BookResponse, error) {
	book := model.Book{}
	if err := bu.br.GetBookByBookId(&book, bookId); err != nil {
		return model.BookResponse{}, err
	}
	bookRes := model.BookResponse(book)
	return bookRes, nil
}

func (bu *bookUsecase) CreateBook(book model.Book) (model.BookResponse, error) {
	if err := bu.br.CreateBook(&book); err != nil {
		return model.BookResponse{}, err
	}
	bookRes := model.BookResponse(book)
	return bookRes, nil
}

func (bu *bookUsecase) UpdateBook(book model.Book, bookId string) (model.BookResponse, error) {
	if err := bu.br.UpdateBook(&book, bookId); err != nil {
		return model.BookResponse{}, err
	}
	bookRes := model.BookResponse(book)
	return bookRes, nil
}

func (bu *bookUsecase) DeleteBook(bookId string) error {
	if err := bu.br.DeleteBook(bookId); err != nil {
		return err
	}
	return nil
}
