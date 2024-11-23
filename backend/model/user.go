package model

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    string `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
