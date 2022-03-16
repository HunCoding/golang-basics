package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type User struct {
	Name string
	Age  int
}

func main() {

	var err error

	db, err = sql.Open("mysql", "root:test123@tcp(172.17.0.2:3306)/users_db")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "HunCoding 2",
		Age:  50,
	}

	if insertError := insertUser(user); insertError != nil {
		panic(err)
	}

	users, err := getAllUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(*user)
	}
}

func getAllUsers() ([]*User, error) {
	res, err := db.Query("SELECT * FROM user_data")
	if err != nil {
		return nil, err
	}

	users := []*User{}

	for res.Next() {

		var user User

		if err := res.Scan(&user.Name, &user.Age); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func insertUser(user User) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO user_data VALUES('%s', %d)", user.Name, user.Age))
	if err != nil {
		return err
	}

	fmt.Println("Usuario inserido com sucesso")
	return nil
}
