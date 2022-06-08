package main

import (
	"fmt"
	"math/rand"
)

type RandomNumber = int64

func main() {
	var randomNumber RandomNumber
	fmt.Println("Estou debuggando codigo em Go!")
	fmt.Printf("Valor atual da variavel: %d \n", randomNumber)

	randomNumber = 20
	addNumber(&randomNumber)

	fmt.Printf("Novo valor da variavel: %d \n", randomNumber)
}

func addNumber(randomNumber *RandomNumber) {
	randNumber := int64(rand.Intn(1000))

	*randomNumber += randNumber
}
