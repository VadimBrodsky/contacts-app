package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	contact "web1.0_app/models"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/contacts")
	})

	e.GET("/contacts", func(c echo.Context) error {
		search := c.QueryParam("q")

		if search != "" {
			contact.Search(search)
		} else {
			contact.All()
		}

		return c.Render(http.StatusOK, "index", "yo")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
