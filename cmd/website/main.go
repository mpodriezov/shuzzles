package main

import (
	"net/http"
	"os"

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
					c.Set("isAuthenticated", true)
				}
			}
			return next(c)
		}
	})

	e.Static("/static", "public")
	e.GET("/", handlers.HandleHomePage)
	e.GET("/404", handlers.HandleNotFound)
	e.GET("/logout", handlers.HandleUserLogout)
	e.GET("/login", handlers.HandleShowUserLogin)
	e.POST("/login", handlers.HandleUserLogin)
	e.GET("/register", handlers.HandleShowUserRegistration)
	e.POST("/register", handlers.HandleNewUserRegistration)
	e.GET("/register/success", handlers.HandlerShowRegistrationSucccess)
	e.POST("/user/:userId/listing", handlers.HandleCreatingNewListing, IsAuthenticated)
	e.GET("/user/:userId/listing", handlers.HandleShowUserListings, IsAuthenticated)

	e.GET("/admin/user", handlers.HandleShowUsersAdminView, IsAuthenticated)
	e.GET("/admin/user/:userId", handlers.HandleShowUserAdminView, IsAuthenticated)
	e.POST("/admin/user/:userId", handlers.HandlerUserAdminUpdate, IsAuthenticated)
	e.GET("/admin/user/:userId/delete", handlers.HandleAdminDeleteUserShow, IsAuthenticated)
	e.DELETE("/admin/user/:userId", handlers.HandleAdminDeleteUser, IsAuthenticated)

	port := os.Getenv("PORT")

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(port))
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ok := c.Get("isAuthenticated").(bool)
		if !ok {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}
