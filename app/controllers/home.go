package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// HomeController : controller
type HomeController struct{}

// Pong : object
type Pong struct {
	Message string `json:"message" xml:"message"`
}

// HomeIndex : path /
func HomeIndex(c echo.Context) error {
	u := &Pong{
		Message: "pong",
	}
	return c.JSON(http.StatusOK, u)
}

// HomeEcho : path /echo
func HomeEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// HomeHello : path /home
func HomeHello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}
