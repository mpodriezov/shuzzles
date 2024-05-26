package main

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/mpodriezov/shuzzles/src/data"
	"github.com/mpodriezov/shuzzles/src/handlers"
	"github.com/mpodriezov/shuzzles/src/templates"
	"github.com/mpodriezov/shuzzles/src/utils"
)

func main() {

	utils.LoadEnvVariables()

	db := data.ConnectSQLDb()
	defer db.Close()

	e := echo.New()
	e.Renderer = templates.CreateRenderer()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dal", &data.Dal{DB: db})
			return next(c)
		}
	})

	// get usr from cookie
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("sid")
			if err == nil {
				// get dal from context
				dal := c.Get("dal").(*data.Dal)
				user := dal.FindUserBySession(cookie.Value)
				if user != nil {
					c.Set("user", user)
				}
			}
			return next(c)
		}
	})

	e.Static("/static", "public")
	e.GET("/", handlers.HandleHomePage)
	e.GET("/login", handlers.HandleShowUserLogin)
	e.POST("/login", handlers.HandleUserLogin)
	e.GET("/register", handlers.HandleShowUserRegistration)
	e.POST("/register", handlers.HandleNewUserRegistration)
	e.GET("/register/success", handlers.HandlerShowRegistrationSucccess)
	e.GET("/protected", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", nil)
	})

	port := os.Getenv("PORT")

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(port))
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		user, ok := sess.Values["user"].(data.SessionUser)
		if !ok {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		c.Set("user", user)
		return next(c)
	}
}
