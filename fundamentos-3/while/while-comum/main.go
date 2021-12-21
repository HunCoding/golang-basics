package main

import "fmt"

func main() {

	// Definindo o valor inicial para comecar o codigo do for com valor 0
	testValue := 0

	// Enquanto testValue for menor ou igual a 20, ele vai executar o codigo que estiver dentro das chaves
	// Similar a outras linguagens nas quais voce faz:
	/*

		while (testValue <= 20) {

		}

	*/
	for testValue <= 20 {
		fmt.Println(fmt.Sprintf("Valor atual: %d", testValue))

		// Adiciona 1 toda vez que passa pelo codigo, para assim, nao entrar em loop infinito
		testValue++
	}

}
