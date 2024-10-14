package model

type Categories struct {
	CategoryID          uint   `json:"category_id" gorm:"primaryKey"`
	CategoryName        string `json:"category_name" gorm:"not null;colum:category_name;size:255"`
	CategoryDescription string `json:"category_description" gorm:"not null;colum:category_description;size:255"`
}
