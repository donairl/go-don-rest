package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func About(c echo.Context) error {
	return c.String(http.StatusOK, "About me !\n")
}

func Home(c echo.Context) error {
	list := [6]string{"Indonesia", "Malaysia", "Vietnam", "Zambia", "Nigeria", "Palestina"}
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Donny",
		"msg":  "Website of root",
		"list": list,
	})

}
