package model

import (
	"time"
)

type Task struct {
	taskName    string
	description string    `bson:"description,omitempty"`
	owner       string    `bson:"owner,omitempty"`
	project     string    `bson:"project,omitempty"`
	dueDate     time.Time `bson:"dueDate,omitempty"`
	completed   bool
}
