package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"regexp"
)

//goland:noinspection SpellCheckingInspection
var DBcon *mongo.Client

type Task struct {
	TaskName    string
	Description string `bson:"description,omitempty"`
	Owner       string `bson:"owner,omitempty"`
	Project     string `bson:"project,omitempty"`
	// need to add date field
	Completed bool   `bson:"completed"`
	ID        string `bson:"_id"`
}

func newTask(taskName, description, owner, project string, completed bool, ID string) Task {
	// process id object to be string
	re := regexp.MustCompile(`(ObjectID|\(|\))`)
	parsedStringId := re.ReplaceAllString(ID, "")
	fmt.Printf("ID: %s; parsedStringID: %s\n", ID, parsedStringId)
	return Task{
		TaskName:    taskName,
		Description: description,
		Owner:       owner,
		Project:     project,
		Completed:   completed,
		ID:          parsedStringId,
	}
}

type Tasks []Task

//func addData() {
//	db := DBcon.Database("GoDo")
//	coll = NewCollection(db)
//	//
//	//task := newTask(
//	//	"do stuff",
//	//	"Go do some stuff",
//	//	"Hunter",
//	//	"Secret project",
//	//)
//
//	res, err := coll.Collection.InsertOne(context.TODO(), task)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Inserted a single document: %v\n", res)
//}

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
	//var returnTasks Tasks
	for _, val := range tasks {
		res = append(res, newTask(val.TaskName, val.Description, val.Owner, val.Project, val.Completed, val.ID))
	}

	return res
}

// TODO - add find one task by id method

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
