package model

import (
	"time"
)

type Book struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Genre       string    `json:"genre"`
	TotalPage   int       `json:"total_page"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishedAt int       `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookResponse struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"not null"`
	Genre       string    `json:"genre"`
	TotalPage   int       `json:"total_page"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishedAt int       `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
