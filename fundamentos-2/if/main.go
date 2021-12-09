package main

import "fmt"

func main() {

	test := "test"

	if test == "test" {
		fmt.Println("Caiu no if")
	} else if test == "test2" {
		fmt.Println("Caiu no else if")
	} else {
		fmt.Println("Caiu no else")
	}
}
