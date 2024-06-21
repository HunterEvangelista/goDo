package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"godo/controller"
	"godo/model"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("MDB_URI")
	if uri == "" {
		log.Fatal("No connection string found.")
	}

	model.DBcon, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = model.DBcon.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	defer func() {
		if err = model.DBcon.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = controller.NewTemplate()
	e.Static("/", "public")

	e.GET("/", controller.HomeRoute)

	//e.POST("/contacts", func(c echo.Context) error {
	//name := c.FormValue("name")
	//email := c.FormValue("email")
	//
	//if page.Data.hasEmail(email) {
	//	formData := newFormData()
	//	formData.Values["name"] = name
	//	formData.Values["email"] = email
	//	formData.Errors["email"] = "Email already exists"
	//	return c.Render(422, "form", formData)
	//}
	//contact := newContact(name, email, id)
	//page.Data.Contacts = append(page.Data.Contacts, contact)
	//err := c.Render(200, "form", newFormData())
	//if err != nil {
	//	return err
	//}
	//return c.Render(200, "oob-contact", contact)
	//})

	//e.DELETE("/contacts/:id", func(c echo.Context) error {
	//	idStr := c.Param("id")
	//	id, err := strconv.Atoi(idStr)
	//	if err != nil {
	//		return c.String(400, "Invalid id")
	//	}
	//
	//	index := page.Data.indexOf(id)
	//	if index == -1 {
	//		return c.String(400, "Contact not found")
	//	}
	//	page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)
	//	return c.NoContent(200)
	//})

	if err = e.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
