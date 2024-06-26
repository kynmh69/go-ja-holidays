package template

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	logger := c.Logger()
	logger.Debug(w, name, data)
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate(path string) *Template {
	return &Template{templates: template.Must(template.ParseGlob(path))}
}
