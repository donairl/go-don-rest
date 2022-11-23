package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code   string `gorm:"type:varchar(8)"`
	Name   string `gorm:"type:varchar(24)"`
	Unit   string `gorm:"type:varchar(12)"`
	Weight float32
	Price  float32
	CategoryID int 
	Category Category
}

type Category struct {
	gorm.Model
	Name string  `gorm:"type:varchar(20)"`
}