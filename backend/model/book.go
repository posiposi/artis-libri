package model

import "time"

type Book struct {
	ID               uint      `json:"id" gorm:"primary_key"`
	Title            string    `json:"title" gorm:"not null"`
	Genre            string    `json:"genre"`
	PageCount        int       `json:"page_count"`
	Author           string    `json:"author"`
	Publisher        string    `json:"publisher"`
	FirstEditionYear int       `json:"first_edition_year"`
	Impression       string    `json:"impression"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	User             User      `json:"user" gorm:"foreignkey:UserID; constraint:onDelete:CASCADE"`
	UserID           uint      `json:"user_id" gorm:"not null"`
}

type BookResponse struct {
	ID               uint      `json:"id" gorm:"primary_key"`
	Title            string    `json:"title" gorm:"not null"`
	Genre            string    `json:"genre"`
	PageCount        int       `json:"page_count"`
	Author           string    `json:"author"`
	Publisher        string    `json:"publisher"`
	FirstEditionYear int       `json:"first_edition_year"`
	Impression       string    `json:"impression"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
