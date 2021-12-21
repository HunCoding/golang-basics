package main

import "fmt"

func main() {

	anExpression := false

	//Aqui o codigo vai passar uma vez e vai sair, pois como e um do-while
	//ele vai executar uma vez, validar, ver que a validacao esta incorreta
	//e assim ele vai sair do for
	for ok := true; ok; ok = anExpression {
		fmt.Println("Passou aqui mesmo com validacao incorreta")
	}

}
