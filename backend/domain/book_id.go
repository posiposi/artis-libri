package domain

import (
	"errors"
)

const BOOK_ID_MAX_LENGTH = 36
const BOOK_ID_ERROR_TEXT = "書籍IDの生成に失敗しました"

type BookId struct {
	value string
}

func NewBookId(value string) (BookId, error) {
	if value == "" {
		return BookId{}, errors.New(BOOK_ID_ERROR_TEXT)
	}
	if len(value) > BOOK_ID_MAX_LENGTH {
		return BookId{}, errors.New(BOOK_ID_ERROR_TEXT)
	}
	return BookId{value: value}, nil
}

func (bi BookId) Value() string {
	return bi.value
}
