package model

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type Book struct {
	ID                 string          `json:"id" gorm:"primary_key"`
	Title              string          `json:"title" gorm:"not null"`
	Genre              string          `json:"genre"`
	TotalPage          int             `json:"total_page"`
	ProgressPage       int             `json:"progress_page"`
	ProgressPercentage int             `json:"progress_percentage"`
	Price              decimal.Decimal `json:"price"`
	Author             string          `json:"author"`
	Publisher          string          `json:"publisher"`
	PublishedAt        int             `json:"published_at"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

func NewBook(
	id, title, genre, author, publisher string,
	totalPage, progressPage, publishedAt int,
	price decimal.Decimal,
) (_ *Book, err error) {
	if title == "" {
		return nil, fmt.Errorf("title is required")
	}
	return &Book{
		ID:                 id,
		Title:              title,
		Genre:              genre,
		TotalPage:          totalPage,
		ProgressPage:       progressPage,
		ProgressPercentage: calcProgressPercentage(totalPage, progressPage),
		Price:              price,
		Author:             author,
		Publisher:          publisher,
		PublishedAt:        publishedAt,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}, nil
}

// TODO レスポンスパッケージに移管しても良いかも
type BookResponse struct {
	ID                 string          `json:"id" gorm:"primary_key"`
	Title              string          `json:"title" gorm:"not null"`
	Genre              string          `json:"genre"`
	TotalPage          int             `json:"total_page"`
	ProgressPage       int             `json:"progress_page"`
	ProgressPercentage int             `json:"progress_percentage"`
	Price              decimal.Decimal `json:"price"`
	Author             string          `json:"author"`
	Publisher          string          `json:"publisher"`
	PublishedAt        int             `json:"published_at"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

func calcProgressPercentage(totalPage, progressPage int) int {
	if totalPage == 0 || progressPage == 0 {
		return 0
	}
	percetange := float64(progressPage) / float64(totalPage) * 100
	return int(percetange)
}
