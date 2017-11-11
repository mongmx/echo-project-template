package Controllers

import (
	"net/http"

	"github.com/labstack/echo"
	models "github.com/mongmx/echo-project-template/app/Models"
)

// HomeController : controller
type HomeController struct{}

// Index : path /
func (HomeController) Index(c echo.Context) error {
	u := &models.Pong{
		Message: "pong",
	}
	return c.JSON(http.StatusOK, u)
}

// Echo : path /echo
func (HomeController) Echo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Hello : path /home
func (HomeController) Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}
