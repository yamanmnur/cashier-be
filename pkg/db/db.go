package db

import (
	"cashier-be/pkg/models"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gorm.DB {

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.TransactionItem{})
	db.AutoMigrate(&models.Transaction{})

	return db
}
