package controller

import (
	"github.com/labstack/echo/v4"
	"godo/model"
	"html/template"
	"io"
	"net/http"
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

// UpdateTask responds with a form for the user to edit the selected task
func UpdateTask(c echo.Context) error {
	task, err := page.Tasks.GetByDisplayID(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.Render(http.StatusOK, "taskUpdate", task)
}

func DeleteTaskRoute(c echo.Context) error {
	_, err := model.DeleteTask(c.Param("id"))
	if err != nil {
		return c.String(400, "Task not found")
	}
	return c.NoContent(200)
}
