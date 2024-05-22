package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserLoginData struct {
	Username string
	password string
}

type UserLoginPage struct {
	Data   UserLoginData
	Errors map[string]string
}

func NewUserLoginPage(username, password string) *UserLoginPage {
	return &UserLoginPage{
		Data: UserLoginData{
			Username: username,
			password: password,
		},
		Errors: nil,
	}
}

func HandleShowUserLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "user/login.html", nil)
}

func HandleUserLogin(c echo.Context) error {

	// username := c.FormValue("username")
	// password := c.FormValue("password")

	// page := NewUserLoginPage(username, password)

	// if !page.Validate() {
	// 	// return 422 Unprocessable Entity along with the page to show the errors
	// 	return c.Render(http.StatusUnprocessableEntity, "user/registration_form.html", page)
	// }

	// check if email or username already exists in database
	//TODO save the user in the database

	return c.Redirect(http.StatusSeeOther, "/")
}
