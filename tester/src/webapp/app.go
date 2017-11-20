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
	log.Printf("index")
	return c.Render(http.StatusOK, "index", map[string]interface{}{})
}

func getSubA(c echo.Context) error {
	log.Printf("sub_a")
	return c.Render(http.StatusOK, "sub_a", map[string]interface{}{})
}

func getSubB(c echo.Context) error {
	log.Printf("sub_b")
	return c.Render(http.StatusOK, "sub_b", map[string]interface{}{})
}

func getSubC(c echo.Context) error {
	log.Printf("sub_c")
	return c.Render(http.StatusOK, "sub_c", map[string]interface{}{})
}

func main() {
	e := echo.New()
	e.Renderer = &Renderer{
		templates: template.Must(template.New("").ParseGlob("views/*.html")),
	}

	e.GET("/", getIndex)
	e.GET("/sub_a", getSubA)
	e.GET("/sub_b", getSubB)
	e.GET("/sub_c", getSubC)

	e.Start(":5000")
}
