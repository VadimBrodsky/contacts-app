package main

import (
	"errors"
	"html/template"
	"io"
	"net/http"

	contact "github.com/VadimBrodsky/contacts-app/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")

	templates := make(map[string]*template.Template)
	templates["contacts.html"] = template.Must(template.ParseFiles("views/contacts.html", "views/base.html"))
	templates["new.html"] = template.Must(template.ParseFiles("views/new.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

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

		return c.Render(http.StatusOK, "contacts.html", data)
	})

	e.GET("/contacts/new", func(c echo.Context) error {
		data := struct {
			Email  string
			Errors map[string]string
		}{Email: "", Errors: make(map[string]string)}
		return c.Render(http.StatusOK, "new.html", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
