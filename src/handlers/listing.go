package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// [GET] user/:id/listing

func HandleShowUserListings(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

// [POST] /user/:id/listing
func HandleCreatingNewListing(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}
