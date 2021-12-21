package main

import "fmt"

func main() {

	// Executando funcao anonima dentro de outra funcao
	func(valueString string, valueInt int) {
		fmt.Println(valueString, valueInt)
	}("TEST TEST", 300)

	// Executando funcao anonima dentro de outra funcao
	func() {
		fmt.Println("EXECUTANDO DE DENTRO DA FUNCAO ANONIMA")
	}()
}
