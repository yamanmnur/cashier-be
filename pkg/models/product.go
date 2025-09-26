package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string  `gorm:"type:varchar(200);not null"`
	Name        string  `gorm:"type:varchar(200);not null"`
	Description string  `gorm:"type:text;not null"`
	Barcode     string  `gorm:"type:varchar(200);not null"`
	Price       float64 `gorm:"type:decimal(20,2);not null"`
	Status      string  `gorm:"type:varchar(20);not null"`
}
