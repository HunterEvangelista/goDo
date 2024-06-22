package model

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var DBcon *mongo.Client

type Task struct {
	TaskName    string
	Description string `bson:"description,omitempty"`
	Owner       string `bson:"owner,omitempty"`
	Project     string `bson:"project,omitempty"`
	// TODO - add due date field
	Completed bool               `bson:"completed"`
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DisplayID string             `bson:"displayId,omitempty"`
}

type Tasks []Task

func newTask(taskName, description, owner, project string, completed bool, ID primitive.ObjectID) Task {
	re := regexp.MustCompile(`(ObjectID|\(|\)|[\!\-\&\;\:\.\,\#\"\']*)`)
	stID := primitive.ObjectID.String(ID)
	parsedStringId := re.ReplaceAllString(stID, "")
	fmt.Println(stID)
	fmt.Println(parsedStringId)
	return Task{
		TaskName:    taskName,
		Description: description,
		Owner:       owner,
		Project:     project,
		Completed:   completed,
		ID:          ID,
		DisplayID:   parsedStringId,
	}
}

func GetTasks() Tasks {
	db := DBcon.Database("GoDo")
	coll := db.Collection("tasks")
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panicf("Error getting tasks: %s\n", err)
	}
	var tasks Tasks
	err = cursor.All(context.TODO(), &tasks)
	if err != nil {
		log.Panicf("Error getting tasks: %s\n", err)
	}
	var res Tasks
	for _, val := range tasks {
		res = append(res, newTask(val.TaskName, val.Description, val.Owner, val.Project, val.Completed, val.ID))
	}

	return res
}

func AddTask(c echo.Context) (Task, error) {
	db := DBcon.Database("GoDo")
	coll := db.Collection("tasks")
	taskname := c.FormValue("taskname")
	desc := c.FormValue("description")
	owner := c.FormValue("owner")
	project := c.FormValue("project")
	completed := c.FormValue("completed")
	completedBool, _ := strconv.ParseBool(completed)
	id := primitive.NewObjectID()
	task := newTask(taskname, desc, owner, project, completedBool, id)
	_, err := coll.InsertOne(context.TODO(), task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

// TODO - add update task function

func DeleteTask(id string) (string, error) {
	db := DBcon.Database("GoDo")
	coll := db.Collection("tasks")
	oId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oId}}
	_, err = coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return "", err
	}
	return "Task deleted", nil
}
