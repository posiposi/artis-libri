package domain

type Review struct {
	reviewId      ReviewId
	reviewContent ReviewContent
}

func NewReview(ri ReviewId, rc ReviewContent) Review {
	return Review{
		reviewId:      ri,
		reviewContent: rc,
	}
}

func (r Review) ReviewId() ReviewId {
	return r.reviewId
}

func (r Review) ReviewContent() ReviewContent {
	return r.reviewContent
}
