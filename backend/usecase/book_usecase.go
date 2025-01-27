package usecase

import (
	"github.com/posiposi/project/backend/infrastructure/openai"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/repository"
)

type IBookUsecase interface {
	GetAllBooks() ([]model.Book, error)
	GetBookByBookId(bookId string) (model.BookResponse, error)
	CreateBook(book model.Book) (model.BookResponse, error)
	UpdateBook(book model.Book, bookId string) (model.BookResponse, error)
	DeleteBook(bookId string) error
	FetchRecommendBooks() (*openai.ChatResponse, error)
}

type bookUsecase struct {
	br repository.IBookRepository
	oc openai.OpenAICommunicator
}

func NewBookUsecase(br repository.IBookRepository, oc openai.OpenAICommunicator) IBookUsecase {
	return &bookUsecase{br, oc}
}

func (bu *bookUsecase) GetAllBooks() ([]model.Book, error) {
	books := []model.Book{}
	if err := bu.br.GetAllBooks(&books); err != nil {
		return nil, err
	}
	bookRes := []model.Book{}
	for _, v := range books {
		t, err := model.NewBook(
			v.ID,
			v.Title,
			v.Genre,
			v.Author,
			v.Publisher,
			v.TotalPage,
			v.ProgressPage,
			v.PublishedAt,
			v.Price,
		)
		if err != nil {
			return nil, err
		}
		bookRes = append(bookRes, *t)
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

func (bu *bookUsecase) FetchRecommendBooks() (*openai.ChatResponse, error) {
	books := []model.Book{}
	repository := bu.br.GetAllBooks(&books)
	if repository != nil {
		return nil, repository
	}

	systemPrompt := &openai.Prompt{
		Role:    openai.System,
		Content: "名古屋の魅力を教えて下さい。",
	}

	var prompts []*openai.Prompt
	prompts = append(prompts, systemPrompt)
	r, err := bu.oc.SendBooks(prompts)
	if err != nil {
		return nil, err
	}
	return r, nil
}
