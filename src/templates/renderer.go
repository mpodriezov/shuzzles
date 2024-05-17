package templates

import (
	"html/template"
	"io"
	"strings"

	"github.com/dannyvankooten/extemplate"
	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *extemplate.Extemplate
}

func (t *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func CreateRenderer() *Renderer {
	xt := extemplate.New().Funcs(template.FuncMap{
		"tolower": strings.ToLower,
	})
	xt.ParseDir("views/", []string{".html"})
	return &Renderer{
		templates: xt,
	}
}
