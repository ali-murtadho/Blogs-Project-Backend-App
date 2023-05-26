package models

import (
	"time"
)

type Blog struct {
	ID          uint   		`gorm:"primary_key" json:"id"`
	Title       string 		`json:"title"`
	Text        string 		`json:"text"`
	// Image       string 	`json:"image"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt	time.Time 	`json:"updated_at"`
	CategoryID  uint 		`json:"categoryid"`
	Category 	Category 	`json:"-"`
}