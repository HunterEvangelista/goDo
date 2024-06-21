package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//goland:noinspection SpellCheckingInspection
var DBcon *mongo.Client

// Collection Set up reference to Tasks collection
type Collection struct {
	Collection *mongo.Collection
}

func NewCollection(db *mongo.Database) *Collection {
	return &Collection{db.Collection("tasks")}
}

var coll *Collection

type Task struct {
	TaskName    string
	Description string `bson:"description,omitempty"`
	Owner       string `bson:"owner,omitempty"`
	Project     string `bson:"project,omitempty"`
	// need to add date field
	Completed bool
	ID        primitive.ObjectID `bson:"_id"`
}

func newTask(taskName, description, owner, project string, completed bool, ID primitive.ObjectID) Task {
	return Task{
		TaskName:    taskName,
		Description: description,
		Owner:       owner,
		Project:     project,
		Completed:   completed,
		ID:          ID,
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
	coll = NewCollection(db)

	cursor, err := coll.Collection.Find(context.TODO(), bson.D{})
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

// TODO - add delete task function
