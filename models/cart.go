package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID int
    User User
	Total int
	Price  float32
	ProductID int 
	Product Product
}

