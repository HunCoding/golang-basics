package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

type User struct {
	ID   string
	Name string
	Age  int
}

func main() {

	var err error

	cluster := gocql.NewCluster("172.18.0.2")
	cluster.Keyspace = "users_keyspace"
	session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	id, err := gocql.RandomUUID()
	if err != nil {
		panic(fmt.Sprintf("Error trying to generate a random uuid, error=%v", err))
	}

	user := User{
		ID:   id.String(),
		Name: "Huncoding",
		Age:  30,
	}

	userInsertResult, err := insertUser(user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted id: %v \n", userInsertResult)

	userGetResult, err := getUserById(user.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resulted user: %v \n", userGetResult)

}

func insertUser(user User) (*User, error) {
	if err := session.Query(
		"INSERT INTO users(user_id, age, name) VALUES (?, ?, ?);",
		user.ID,
		user.Age,
		user.Name,
	).Exec(); err != nil {
		return nil, err
	}

	return &user, nil
}

func getUserById(id string) (*User, error) {

	var user User

	if err := session.Query(
		"SELECT user_id, age, name FROM users WHERE user_id = ?",
		id).Scan(
		&user.ID,
		&user.Age,
		&user.Name,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
