package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		if err := godotenv.Load("/root/apps/.home.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	serverIp := os.Getenv("SERVER_IP")
	port := os.Getenv("PORT")
	ssl, _ := strconv.ParseBool(os.Getenv("SSL"))

	e := echo.New()
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("rustydoggobytes.com")
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		component := layout(index())
		return component.Render(c.Request().Context(), c.Response())
	})

	serverAddress := serverIp + ":" + port
	if ssl {
		e.Logger.Fatal(e.StartAutoTLS(serverAddress))
	} else {
		e.Logger.Fatal(e.Start(serverAddress))
	}
}
