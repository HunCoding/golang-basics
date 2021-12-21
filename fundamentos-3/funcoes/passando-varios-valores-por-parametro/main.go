package main

import "fmt"

func main() {

	receivingStringParameters("TEST", "TEST", "TEST")

	funcaoTest := func() {
		fmt.Println("TEST DE DENTRO DA FUNCAO")
	}

	receivingFunctionParameters(funcaoTest, funcaoTest)

	funcaoTestWithParameters := func(valueString string, valueInt int) {
		fmt.Println(valueInt, valueString)
	}

	receivingFunctionWithParameterByParameters(funcaoTestWithParameters, funcaoTestWithParameters)

}

func receivingStringParameters(stringsValues ...string) {
	for _, x := range stringsValues {
		fmt.Println(x)
	}
}

func receivingFunctionParameters(stringsValues ...func()) {
	for _, x := range stringsValues {
		x()
	}
}

func receivingFunctionWithParameterByParameters(stringsValues ...func(string, int)) {
	for _, x := range stringsValues {
		x("TEST", 300)
	}
}
