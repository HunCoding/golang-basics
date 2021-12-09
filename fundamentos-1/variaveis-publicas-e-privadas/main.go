package main

import "fmt"

var (
	TestPublica = "test"
	testPrivada = "test"
)

func main() {

	/*

			 Como estamos no mesmo pacote, mesmo sendo privada ainda assim é acessivel aqui dentro
		   porem, ao criar novos pacotes, a variavel testPrivada e a funcao "funcaoPrivada" não
		   será visivel

	*/

	funcaoPrivada()
	FuncaoPublica()

	fmt.Println(TestPublica)
	fmt.Println(testPrivada)

}

func funcaoPrivada() {
	return
}

func FuncaoPublica() {

}
