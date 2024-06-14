package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var DB *mongo.Client

func Db() (*mongo.Client, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("MDB_URI")
	if uri == "" {
		log.Fatal("MDB_URI env variable not set")
	}
	DB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := DB.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	err = DB.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	return DB, nil
}
