package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mpodriezov/shuzzles/src/data"
	"github.com/mpodriezov/shuzzles/src/utils"
)

const SESSION_TIME_HOURS = 24

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
		Errors: map[string]string{},
	}
}

func HandleShowUserLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "user/login.html", nil)
}

func HandleUserLogin(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	page := NewUserLoginPage(username, "")

	if username == "" || password == "" {
		page.Errors["error"] = "Username and password are required"
		return c.Render(http.StatusUnprocessableEntity, "user/login_form.html", page)
	}

	dal := c.Get("dal").(*data.Dal)
	user, err := dal.FindUserByUsername(username)

	if err != nil {
		panic(err)
	}

	if user == nil {
		c.Logger().Error("User not found " + username)
		page.Errors["error"] = "Invalid username or password"
		return c.Render(http.StatusUnprocessableEntity, "user/login_form.html", page)
	}

	if !utils.CheckHash(password, user.PasswordHash) {
		c.Logger().Error("Invalid password for user " + username)
		page.Errors["error"] = "Invalid username or password"
		return c.Render(http.StatusUnprocessableEntity, "user/login_form.html", page)
	}

	sessionID, err := utils.GenerateSessionID()
	if err != nil {
		c.Logger().Error(err)
		page.Errors["error"] = "Unknown error occurred. Please try again later."
		return c.Render(http.StatusUnprocessableEntity, "user/login_form.html", page)
	}

	expiresAt := time.Now().Add(time.Duration(SESSION_TIME_HOURS) * time.Hour)

	//create session in db
	session := dal.CreateSession(sessionID, user.Id, expiresAt)
	// create session cookie
	utils.CreateSessionCookie(c, session.SessionId, expiresAt)

	return c.Redirect(http.StatusSeeOther, "/")
}

func HandleUserLogout(c echo.Context) error {
	sessionCookie, err := c.Cookie("sid")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	dal := c.Get("dal").(*data.Dal)
	dal.DeleteSession(sessionCookie.Value)
	utils.DeleteSessionCookie(c)
	return c.Redirect(http.StatusSeeOther, "/login")
}
