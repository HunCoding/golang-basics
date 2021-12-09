package main

import "fmt"

func main() {
	var testSlice []string = []string{}
	var testArray [2]string = [2]string{"test", "test2"}

	fmt.Println(testSlice, testArray)

	// Printando a capacidade e o tamanho atual do
	// slice, para em seguida, inserir mais um valor
	// para comparar os novos tamanhos
	fmt.Println(cap(testSlice))
	fmt.Println(len(testSlice))

	testSlice = append(testSlice, "testValorAppend")

	fmt.Println(cap(testSlice))
	fmt.Println(len(testSlice))
}
