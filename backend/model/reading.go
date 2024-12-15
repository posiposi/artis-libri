package model

type Reading struct {
	UserId       string `json:"userId" gorm:"primary_key"`
	BookId       string `json:"bookId" gorm:"primary_key"`
	ProgressPage int    `json:"progressPage"`
	Review       string `json:"review"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
