package main

import "fmt"

func main() {

	//Definindo a lista na qual o for vai iterar
	listOfStringValues := []string{"test", "test", "test"}

	// O range retorna dois valores, o primeiro deles sendo o indice
	// e o segundo sendo o valor em si dentro da lista na qual voce
	// esta iterando
	for i, value := range listOfStringValues {
		fmt.Println(fmt.Sprintf("Valor: %s, indice: %d", value, i))
	}
}
