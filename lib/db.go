package lib

import (
	"github.com/donairl/go-don-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=ecommerce_dev port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
   
	db.AutoMigrate(&models.User{},&models.Product{},&models.Category{}, &models.Cart{})

	return db
}
