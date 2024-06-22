package controller

import (
	"godo/model"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

var page Page

func HomeRoute(c echo.Context) error {
	page = NewPage()
	return c.Render(http.StatusOK, "index", page)
}

func CreateTask(c echo.Context) error {
	task, err := model.AddTask(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	page.Tasks = append(page.Tasks, task)
	return c.Render(http.StatusOK, "oob-task", task)
}

func DeleteTaskRoute(c echo.Context) error {
	// get id from request and pass to model
	_, err := model.DeleteTask(c.Param("id"))
	if err != nil {
		return c.String(400, "Task not found")
	}
	return c.NoContent(200)
}
