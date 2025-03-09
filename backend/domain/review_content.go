package domain

import "errors"

type ReviewContent struct {
	value string
}

func NewReviewContent(value string) (ReviewContent, error) {
	if value == "" {
		return ReviewContent{}, errors.New("レビュー内容の生成に失敗しました")
	}
	return ReviewContent{value: value}, nil
}

func (rc ReviewContent) Value() string {
	return rc.value
}
