package main

import "fmt"

func assertValue(value1 string, value2 string) {
	if value1 == value2 {
		fmt.Println("Sao iguais")
	}

	panic("Nao sao iguais")
}

func main() {
	try {
		assertValue("hun", "coding")
	} catch(e Exception) {
		fmt.Println
	}
}
