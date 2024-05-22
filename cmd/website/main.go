package main

import (
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

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dal", &data.Dal{DB: db})
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

	port := os.Getenv("PORT")

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(port))
}
