package model

import "time"

type Reading struct {
	Id           int       `json:"id" gorm:"primary_key"`
	UserId       string    `json:"userId" gorm:"primary_key"`
	BookId       string    `json:"bookId" gorm:"primary_key"`
	ProgressPage int       `json:"progressPage"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
