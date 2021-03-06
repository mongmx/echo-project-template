package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	c "github.com/mongmx/echo-project-template/app/Controllers"
)

// Template : template renderer
type Template struct {
	templates *template.Template
}

// Render : render function
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("./app/views/*.html")),
	}
	e.Renderer = t

	e.Static("/public", "public")
	e.File("/favicon.ico", "public/favicon.ico")

	e.GET("/", c.HomeController{}.Index)
	e.GET("/home", c.HomeController{}.Hello)

	e.Logger.Fatal(e.Start(":3001"))
}
