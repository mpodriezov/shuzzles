package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mpodriezov/shuzzles/src/handlers"
	"github.com/mpodriezov/shuzzles/src/templates"
)

func main() {
	e := echo.New()
	e.Renderer = templates.CreateRenderer()
	e.Use(middleware.Logger())
	e.Static("/static", "public")
	e.GET("/", handlers.HandleHomePage)
	e.GET("/register", handlers.HandleUserRegistrationPage)
	e.POST("/register", handlers.HandleUserRegistration)
	e.Logger.Fatal(e.Start(":9999"))
}
