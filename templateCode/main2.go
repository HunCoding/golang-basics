package main

import "fmt"

type Person struct {
	Name string
	Youtube string

}

func printStruct(s Person) {
   fmt.Println(s)
}

func main() {
	s := Person{
				Name: "John",
				Youtube: "youtube.com/@huncoding",
		
	}

	printStruct(s)
}
	