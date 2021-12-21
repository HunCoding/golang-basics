package main

import "fmt"

func main() {

	listOfValues := []string{"test", "test", "test"}

	for i := 0; i < len(listOfValues); i++ {
		fmt.Println(fmt.Sprintf("Valor atual: %s, indice do valor: %d", listOfValues[i], i))
	}
}
