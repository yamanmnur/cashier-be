package models

import "gorm.io/gorm"

type Role string

const (
	ADMIN   Role = "ADMIN"
	CASHIER Role = "CASHIER"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;index"`
	Password string `gorm:"type:varchar(200);not null"`
	Name     string `gorm:"type:varchar(200);not null"`
	Role     Role   `gorm:"column:role" sql:"type:enum('ADMIN','CASHIER');default:'ADMIN'"`
}
