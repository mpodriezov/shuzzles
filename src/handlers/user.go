package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleUserRegistrationPage(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func HandleUserRegistration(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
