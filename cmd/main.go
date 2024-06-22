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

	e.DELETE("/task/:id", controller.DeleteTaskRoute)

	e.POST("/task", controller.CreateTask)

	e.GET("/task/:id", controller.UpdateTask)
	// TODO - create route to process updated task submission
	//e.POST("/task/:id", controller.SubmitUpdate)

	if err = e.Start(":8080"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
