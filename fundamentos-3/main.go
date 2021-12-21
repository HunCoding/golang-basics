package main

import "fmt"

func main() {

	testParametro2 := func() {
		fmt.Println("test2")
	}
	testParametro := func() {
		fmt.Println("test")
	}
	test(testParametro, testParametro2)

}

func test(valoresString ...func()) {

	for _, x := range valoresString {
		x()
	}
}
