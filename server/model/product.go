package model

import "gorm.io/datatypes"

type Products struct {
	ProductID           uint           `json:"product_id" gorm:"primaryKey"`
	ProductName         string         `json:"product_name" gorm:"not null;column:product_name;size:255"`
	ProductSerialNumber string         `json:"product_serial_number" gorm:"not null;column:product_serial_number;size:100"`
	ProductIMG          string         `json:"product_img" gorm:"not null;column:product_img;size:255"`
	AdditionalInfo      datatypes.JSON `json:"additional_info" gorm:"not null"`
	CategoryID          int            `json:"category_id" gorm:"foreignKey:category_id"`
}

type ListProducts struct {
	ProductID           uint           `json:"product_id" gorm:"primaryKey"`
	ProductName         string         `json:"product_name" gorm:"not null;column:product_name;size:255"`
	ProductSerialNumber string         `json:"product_serial_number" gorm:"not null;column:product_serial_number;size:100"`
	ProductIMG          string         `json:"product_img" gorm:"not null;column:product_img;size:255"`
	AdditionalInfo      datatypes.JSON `json:"additional_info" gorm:"not null"`
	CategoryName        string         `json:"category_name" gorm:"not null;column:category_name;size:255"`
}
