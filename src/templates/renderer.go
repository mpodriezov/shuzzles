package templates

import (
	"html/template"
	"io"
	"strings"

	"github.com/dannyvankooten/extemplate"
	"github.com/labstack/echo/v4"
	"github.com/mpodriezov/shuzzles/src/data"
)

type Renderer struct {
	templates *extemplate.Extemplate
}

type TemplateData struct {
	IsAuthenticated bool
	User            any
	Ctx             any
}

func NewTemplateData() TemplateData {
	return TemplateData{
		IsAuthenticated: false,
		User:            data.SessionUser{},
		Ctx:             nil,
	}
}

func (t *Renderer) Render(w io.Writer, name string, ctxData any, c echo.Context) error {
	data := NewTemplateData()
	data.Ctx = ctxData
	user := c.Get("user")
	if user != nil {
		data.User = user
		data.IsAuthenticated = true
	}
	c.Logger().Info(data)
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
