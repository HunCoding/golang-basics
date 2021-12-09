package main

import "fmt"

func main() {

	test := "TEST_TEST"

	switch test {

	case "test", "test2", "test434", "test1323":
		fmt.Print("CAIU NA PRIMEIRA CONDICAO")
		fallthrough

	case "test_case_2":
		fmt.Println("CAIU NA SEGUNDA CONDIÇÃO")
		fallthrough

	case "test_case_3":
		fmt.Println("CAIU NA TERCEIRA CONDIÇÃO")
		break

	default:
		fmt.Println("CAIU NO DEFAULT")
	}

}
