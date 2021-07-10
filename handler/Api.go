package handler

import (
	"net/http"

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
}

func fake(c echo.Context) error {
	myproduct := models.Product{Name: "Shampo A", Code: "10020440", Weight: 21.2, Unit: "pcs"}
	db.Create(&myproduct)
	return c.String(http.StatusOK, "Fake data created !\n")
}

func adminSite(c echo.Context) error {
	return c.String(http.StatusOK, "This is an index in admin Site !\n")
}
