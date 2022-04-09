package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

var (
	client *elasticsearch.Client
)

func main() {

	cert, err := ioutil.ReadFile("/home/hunter/http_ca.crt")
	if err != nil {
		panic(err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		CACert:   cert,
		Username: "elastic",
		Password: "lEUtD-LaqO2sZnTsLwrH",
	}

	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	_, err = client.Info()
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "HunCoding",
		Age:  30,
	}

	insertedId, err := Index(user)
	if err != nil {
		panic(err)
	}

	res, err := GetUser(insertedId)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func Index(user User) (string, error) {

	id := rand.Intn(100)
	idUser := strconv.Itoa(id)

	requestBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	cfg := esapi.IndexRequest{
		DocumentID: idUser,
		Index:      "users",
		Body:       bytes.NewReader(requestBytes),
		Refresh:    "true",
	}

	res, err := cfg.Do(context.Background(), client)
	if err != nil {
		return "", err
	}

	fmt.Println(res)

	return idUser, nil
}

func GetUser(id string) (*esapi.Response, error) {

	cfg := esapi.GetRequest{
		Index:      "users",
		DocumentID: id,
	}

	res, err := cfg.Do(context.Background(), client)
	if err != nil {
		return nil, err
	}

	return res, nil
}
