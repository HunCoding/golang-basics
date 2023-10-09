package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Collection     *mongo.Collection
	CollectionName string = "testUserCol"
	DatabaseName   string = "testUserDb"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Println("Error trying to open connection")
		return
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println("Error trying to open connection")
		return
	}

	Collection = client.Database(DatabaseName).Collection(CollectionName)

	user, err := CreateUser("Huncoding")
	fmt.Println(user)
}

type User struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func CreateUser(
	name string,
) (*User, error) {
	user := bson.M{"name": name}
	result, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	userObj := &User{
		Name: name,
	}

	userObj.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return userObj, nil
}
