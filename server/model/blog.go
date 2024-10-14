package model

type Blog struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null;colum:title;size:255"`
	Post  string `json:"post" gorm:"not null;colum:post;size:255"`
}
