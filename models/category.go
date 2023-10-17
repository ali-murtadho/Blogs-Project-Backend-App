package models

type Category struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	CatType string `json:"cat_type"`
	Blog    []Blog `json:"-"`
}
