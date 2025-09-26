package models

import "gorm.io/gorm"

type Status string

const (
	COMPLETED Status = "COMPLETED"
	CANCEL    Status = "CANCEL"
	REFUND    Status = "REFUND"
	PENDING   Status = "PENDING"
)

type Transaction struct {
	gorm.Model
	InvoiceNo    string            `gorm:"type:varchar(200);not null"`
	UserId       uint              `gorm:"not null"`
	UserName     string            `gorm:"type:varchar(200);not null"`
	CustomerName string            `gorm:"varchar(200);not null"`
	PhoneNumber  string            `gorm:"varchar(30);not null"`
	Total        float64           `gorm:"type:decimal(20,2);not null"`
	Status       Status            `gorm:"column:status" sql:"type:enum('COMPLETED','PENDING','CANCEL', 'REFUND');default:'PENDING'"`
	Items        []TransactionItem `json:"items" gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
