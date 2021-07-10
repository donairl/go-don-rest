package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Site(c echo.Context) error {
	return c.String(http.StatusOK, "This is a Site !\n")
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "This is a Home sweet home 2 !\n")
}
