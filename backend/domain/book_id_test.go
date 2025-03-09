package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBookId(t *testing.T) {
	t.Run("正常系(最大値)", func(t *testing.T) {
		value := "d04a8be984d6d2318e7268bb1b79cfa07890"
		book, _ := NewBookId(value)
		assert.Equal(t, book, BookId{value: value})
	})

	t.Run("書籍IDが存在しない", func(t *testing.T) {
		value := ""
		expected := "書籍IDの生成に失敗しました"
		_, err := NewBookId(value)
		assert.EqualError(t, err, expected)
	})

	t.Run("最大文字列超過(37文字)", func(t *testing.T) {
		value := "d04a8be984d6d2318e7268bb1b79cfa0789071"
		expected := "書籍IDの生成に失敗しました"
		_, err := NewBookId(value)
		assert.EqualError(t, err, expected)
	})
}

func TestBookIdValue(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		value := "value test"
		bookId, _ := NewBookId(value)
		assert.Equal(t, bookId.Value(), value)
	})
}
