package main

import "fmt"

func main() {

	userVar := user{
		name:  "test",
		age:   20,
		test2: "test2",
	}

	testVar := test{}

	fmt.Println(userVar)
	fmt.Println(testVar)

}

type test struct {
}

type user struct {
	name  string
	age   int
	test2 string
}
