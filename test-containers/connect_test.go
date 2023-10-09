package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestMain(m *testing.M) {
	closeConnection := OpenConnection()
	defer closeConnection()
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	testUser, err := CreateUser("test")
	if err != nil {
		t.FailNow()
		return
	}

	if testUser.Name != "test" {
		t.FailNow()
		return
	}

	hex, err := primitive.ObjectIDFromHex(testUser.ID)
	if err != nil {
		return
	}
	filter := bson.D{{Key: "_id", Value: hex}}
	result := Collection.FindOne(context.Background(), filter)
	user := User{}
	err = result.Decode(&user)
	if err != nil {
		t.FailNow()
		return
	}
}
