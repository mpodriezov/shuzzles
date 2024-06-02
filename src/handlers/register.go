package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mpodriezov/shuzzles/src/data"
	"github.com/mpodriezov/shuzzles/src/utils"
)

const MAX_PASSWORD_LENGTH = 8

const RoleAdmin = 2
const RoleUser = 4

type UserRegistraionData struct {
	Username        string
	Email           string
	password        string
	passwordConfirm string
}

type UserRegistrationPage struct {
	Data   UserRegistraionData
	Errors map[string]string
}

func NewUserRegistrationPage(c echo.Context) *UserRegistrationPage {
	return &UserRegistrationPage{
		Data: UserRegistraionData{
			Username:        c.FormValue("username"),
			Email:           c.FormValue("email"),
			password:        c.FormValue("password"),
			passwordConfirm: c.FormValue("confirm_password"),
		},
		Errors: nil,
	}
}

func (p *UserRegistrationPage) Validate() bool {
	p.Errors = make(map[string]string)

	if len(p.Data.Username) < 3 {
		p.Errors["username"] = "Username should be at least 3 characters long"
	}

	if !utils.IsEmailValid(p.Data.Email) {
		p.Errors["email"] = "Invalid email address"
	}

	if p.Data.password != p.Data.passwordConfirm {
		p.Errors["password"] = "Passwords do not match"
	}

	if !utils.IsPasswordComplex(p.Data.password, MAX_PASSWORD_LENGTH, true, true, true, true) {
		p.Errors["password"] = "Password should contain at least one uppercase letter and one lowercase letter"
	}
	return len(p.Errors) == 0
}

// handler for [GET] /register
func HandleShowUserRegistration(c echo.Context) error {
	return c.Render(http.StatusOK, "user/register.html", NewUserRegistrationPage(c))
}

// handler for [POST] /register
func HandleNewUserRegistration(c echo.Context) error {

	page := NewUserRegistrationPage(c)

	if !page.Validate() {
		// return 422 Unprocessable Entity along with the page to show the errors
		return c.Render(http.StatusUnprocessableEntity, "user/registration_form.html", page)
	}

	dal := c.Get("dal").(*data.Dal)

	passwordhash, _ := utils.Hash(page.Data.password)

	usr := data.RegistrationModel{
		Username:     page.Data.Username,
		Email:        page.Data.Email,
		PasswordHash: passwordhash,
		Bio:          "",
		Role:         RoleUser,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	c.Logger().Debug("new user registration data: ", usr)

	exists := dal.UserExists(usr.Username, usr.Email)

	if exists {
		page.Errors["username"] = "Username or email already exists"
		return c.Render(http.StatusUnprocessableEntity, "user/registration_form.html", page)
	}

	userId := dal.NewUser(&usr)

	c.Logger().Info("new user created with id: ", userId)

	c.Response().Header().Set("HX-Redirect", "/register/success")
	return c.NoContent(http.StatusFound)
}

// handler for [GET] register/success
func HandlerShowRegistrationSucccess(c echo.Context) error {
	return c.Render(http.StatusOK, "user/registration_result.html", nil)
}
