package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReview(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		ri, _ := NewReviewId(uint64(1))
		rc, _ := NewReviewContent("とても良い作品")
		review := NewReview(ri, rc)
		assert.Equal(t, review, Review{reviewId: ri, reviewContent: rc})
	})
}

func TestReviewId(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		ri, _ := NewReviewId(uint64(1))
		rc, _ := NewReviewContent("とても良い作品")
		review := NewReview(ri, rc)
		assert.Equal(t, review.ReviewId(), ri)
	})
}

func TestReviewContent(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		ri, _ := NewReviewId(uint64(1))
		rc, _ := NewReviewContent("とても良い作品")
		review := NewReview(ri, rc)
		assert.Equal(t, review.ReviewContent(), rc)
	})
}
