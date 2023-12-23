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

		var contacts []contact.Contact
		if search != "" {
			contacts, _ = contact.Search(search)
		} else {
			contacts, _ = contact.All()
		}

		data := struct {
			Search   string
			Contacts []contact.Contact
		}{Search: search, Contacts: contacts}

		return c.Render(http.StatusOK, "layout.html", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
