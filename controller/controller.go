package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"godo/model"
)

// controller will communicate with pages and model to update and serve data

//type Templates struct {
//	templates *template.Template
//}
//
//func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//	return t.templates.ExecuteTemplate(w, name, data)
//}
//
//func NewTemplate() *Templates {
//	return &Templates{
//		template.Must(template.ParseGlob("views/*")),
//	}
//}

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

// TODO - add back views to air config

func HomeRoute(c echo.Context) error {
	page := NewPage()
	//return c.Render(http.StatusOK, "index", page)
	fmt.Println(page)
	return c.NoContent(200)
}
