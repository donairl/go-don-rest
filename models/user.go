package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(32)"`
	Password    string `gorm:"type:varchar(32)"`
	Name   string `gorm:"type:varchar(60)"`
	Email   string `gorm:"type:varchar(60)"`
	Address   string `gorm:"type:varchar(200)"`
	Postalcode   string `gorm:"type:varchar(20)"`

}

