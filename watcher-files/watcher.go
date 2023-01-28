package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fsnotify/fsnotify"
	yaml "gopkg.in/yaml.v2"
)

var (
	Configurations User
)

type User struct {
	Username string `json:"username" yaml:"username"`
	Lastname string `json:"lastname" yaml:"lastname"`
	Age      int8   `json:"age" yaml:"age"`
}

func ReadAndMarshalFile() {
	fmt.Println("Reading configuration")
	data, err := ioutil.ReadFile("configurations.yml")
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal([]byte(data), &Configurations); err != nil {
		panic(err)
	}

	fmt.Printf("Configuration loaded: %#v \n", Configurations)
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File changed, changing configuration")
					ReadAndMarshalFile()
					fmt.Printf("New configuration: %#v\n", Configurations)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error trying to load file, error:", err)
			}
		}
	}()

	//Adding listener to configuration file
	err = watcher.Add("configurations.yml")
	if err != nil {
		log.Fatal(err)
	}

	//Read and load configuration
	ReadAndMarshalFile()

	//Setting code to wait
	<-done
}
