package model

type Stocks struct {
	StockID      uint `json:"stock_id" gorm:"primaryKey"`
	ProductID    int  `json:"product_id" gorm:"foreignKey:product_id"`
	StockProduct int  `json:"stock_product" gorm:"not null;colum:stock_product"`
}
