package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		component := layout(index())
		return component.Render(c.Request().Context(), c.Response())
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
