package main

import (
	"github.com/labstack/echo/v4"
  	"github.com/labstack/echo/v4/middleware"
	"echo-app/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.DBMigrate()
	initializers.AddLogger()
	initializers.ConnectToCache()
}

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", Hola)

	app.Logger.Fatal(app.Start(":" + initializers.CONFIG.PORT))
}

func Hola(c echo.Context) error {
	return c.String(200, "Hola Amigo!")
  }