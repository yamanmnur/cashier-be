package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint        `gorm:"not null;index"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
	ProductId     string      `gorm:"type:varchar(200);not null"`
	Code          string      `gorm:"type:text;not null"`
	ProductName   string      `gorm:"type:varchar(200);not null"`
	Barcode       string      `gorm:"type:varchar(200);not null"`
	Price         float64     `gorm:"type:decimal(20,2);not null"`
	Quantity      float64     `gorm:"type:decimal(20,2);not null"`
	Total         float64     `gorm:"type:decimal(20,2);not null"`
}
