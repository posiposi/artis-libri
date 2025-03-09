package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReviewContent(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		value := "validated review content"
		review, _ := NewReviewContent(value)
		assert.Equal(t, review, ReviewContent{value: value})
	})

	t.Run("レビューが存在しない", func(t *testing.T) {
		value := ""
		expected := "レビュー内容の生成に失敗しました"
		_, err := NewReviewContent(value)
		assert.EqualError(t, err, expected)
	})
}

func TestReviewContentValue(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		value := "value test"
		review, _ := NewReviewContent(value)
		assert.Equal(t, review.Value(), value)
	})
}
