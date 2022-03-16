package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  int32              `bson:"age"`
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	user := User{Name: "Huncoding",
		Age: 30,
	}

	collection := client.Database("user").Collection("userData")

	//INSERT ONE
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)

	//FIND ONE
	filter := bson.D{{"name", user.Name}}
	userResult := User{}
	errFinding := collection.FindOne(context.Background(), filter).Decode(&userResult)
	if errFinding != nil {
		panic(errFinding)
	}

	fmt.Println(userResult)

	//FIND ALL
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		userResult := User{}

		err := cur.Decode(&userResult)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userResult)
	}
}
