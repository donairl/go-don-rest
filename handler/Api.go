package handler

import (
	"net/http"
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
	g.GET("", adminSite)
	g.GET("/fake", fake)
	g.GET("/products", product)
	e.GET("/products/:id", getUser)
}

func fake(c echo.Context) error {
	myproduct := models.Product{Name: "Shampo A", Code: "10020440", Weight: 21.2, Unit: "pcs"}
	db.Create(&myproduct)
	myproduct1 := models.Product{Name: "Shoap Body", Code: "10020441", Weight: 18, Unit: "pcs"}
	db.Create(&myproduct1)

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
	db.First(&products,id)
	return c.JSON(http.StatusOK, products)
}

func adminSite(c echo.Context) error {
	return c.String(http.StatusOK, "This is an index in admin Site !\n")
}
