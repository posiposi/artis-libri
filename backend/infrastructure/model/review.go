package model

import "time"

type Review struct {
	Id            int       `json:"id" gorm:"primary_key"`
	ReviewContent string    `json:"review_content"`
	Reading       Reading   `json:"reading"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
