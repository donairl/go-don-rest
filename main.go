package main

import (
	"html/template"
	"io"

	"github.com/donairl/go-don-rest/lib"
	"github.com/donairl/go-don-rest/router"
	"github.com/labstack/echo/v4"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates *template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance

	db := lib.ConnectDb()
	e := router.New(db)
	// Instantiate a template registry and register all html files inside the view folder
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
