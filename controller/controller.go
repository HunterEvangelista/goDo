package controller

import (
	"github.com/labstack/echo/v4"
	"godo/model"
	"html/template"
	"io"
	"net/http"
)

//controller will communicate with pages and model to update and serve data

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		template.Must(template.ParseGlob("views/*")),
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Tasks model.Tasks
	Form  FormData
}

func NewPage() Page {
	return Page{
		Tasks: model.GetTasks(),
		Form:  NewFormData(),
	}
}

func HomeRoute(c echo.Context) error {
	page := NewPage()
	return c.Render(http.StatusOK, "index", page)
}

func DeleteTaskRoute(c echo.Context) error {
	// get id from request and pass to model
	_, err := model.DeleteTask(c.Param("id"))
	if err != nil {
		return c.String(400, "Task not found")
	}
	// respond with 200
	return c.NoContent(200)
}
