package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health_check", healthCheck)
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func hello(c echo.Context) error {
	user := c.QueryParam("user")
	return c.String(http.StatusOK, "Hello, "+user)
}
