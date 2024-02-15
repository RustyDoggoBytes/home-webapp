package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	assetHandler := http.FileServer(getFileSystem(false))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	e.GET("/", func(c echo.Context) error {
		component := layout(index())
		return component.Render(c.Request().Context(), c.Response())
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}
