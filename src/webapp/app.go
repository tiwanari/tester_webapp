package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func getIndex(c echo.Context) error {
	log.Printf("init")
	return c.Render(http.StatusOK, "index", map[string]interface{}{})
}

func main() {
	e := echo.New()
	e.Renderer = &Renderer{
		templates: template.Must(template.New("").ParseGlob("views/*.html")),
	}

	e.GET("/", getIndex)

	e.Start(":5000")
}
