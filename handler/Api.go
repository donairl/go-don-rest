package handler

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/donairl/go-don-rest/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterApi(e *echo.Echo, dbGlobal *gorm.DB) {
	g := e.Group("/api")
	g.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin test !\n")
	})
	db = dbGlobal
	g.GET("", apiSite)
	g.GET("/fake", fake)
	g.GET("/products", product)
	e.GET("/products/:id", getProductByCode)

}

func fake(c echo.Context) error {

	//myproduct := models.Product{Name: "Shampo A", Code: "10020440", Weight: 21.2, Unit: "pcs",Price:0}
	//db.Create(&myproduct)

	//2023 Now change read from csv
	f, err := os.Open("products.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range data {
		if i > 0 { // omit header line
			var rec models.Product
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 1 {
					rec.Code = field
				} else if j == 2 {

					wvalue, errw := strconv.ParseFloat(field, 32)
					if err != nil {
						// do something sensible
						log.Fatal(errw)
					}
					rec.Weight = float32(wvalue)

				} else if j == 3 {
					rec.Unit = field

				} else if j == 4 {

					value, err := strconv.ParseFloat(field, 32)
					if err != nil {
						// do something sensible
						log.Fatal(err)
					}
					rec.Price = float32(value)

				}
			}
			rec.CategoryID = 1
			db.Create(&rec)
		}
	}

	return c.String(http.StatusOK, "Fake data created !\n")
}

func product(c echo.Context) error {
	var products []models.Product
	db.Find(&products)

	return c.JSON(http.StatusOK, products)
}

func getProductByCode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var products []models.Product
	db.First(&products, id)
	return c.JSON(http.StatusOK, products)
}

func apiSite(c echo.Context) error {
	return c.String(http.StatusOK, "This is an index in api Site !\n")
}
