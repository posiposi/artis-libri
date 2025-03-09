package domain

import (
	"errors"
)

const REVIEW_ID_ERROR_TEXT = "レビューIDの生成に失敗しました"
const REVIEW_ID_MAX_LENGTH = 18446744073709551615

type ReviewId struct {
	value uint64
}

func NewReviewId(value uint64) (ReviewId, error) {
	if value > REVIEW_ID_MAX_LENGTH {
		return ReviewId{}, errors.New("レビューIDの生成に失敗しました")
	}
	return ReviewId{value: value}, nil
}

func (ri ReviewId) Value() uint64 {
	return ri.value
}
