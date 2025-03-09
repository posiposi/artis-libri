package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReviewId(t *testing.T) {
	t.Run("正常系：最小値", func(t *testing.T) {
		reviewId, _ := NewReviewId(0)
		assert.Equal(t, reviewId, ReviewId{value: 0})
	})

	t.Run("正常系：最大値", func(t *testing.T) {
		reviewId, _ := NewReviewId(18446744073709551615)
		assert.Equal(t, reviewId, ReviewId{value: 18446744073709551615})
	})
}

func TestReviewIdValue(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		value := uint64(1)
		reviewId, _ := NewReviewId(value)
		assert.Equal(t, reviewId.Value(), value)
	})
}
