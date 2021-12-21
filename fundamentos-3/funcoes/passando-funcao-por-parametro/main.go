package main

import "fmt"

func main() {

	funcaoTesteParametro := func(valueString string, valueInt int) {
		fmt.Println(valueString, valueInt)
	}

	testFunctionByParameter(funcaoTesteParametro)

	funcao := returningFunctions()

	funcao("TEST", 4000)

}

func returningFunctions() func(string, int) {
	funcaoTestRetorno := func(valueString string, valueInt int) {
		fmt.Println(valueString, valueInt)
	}

	return funcaoTestRetorno
}

func testFunctionByParameter(funcaoTest func(string, int)) {
	funcaoTest("TEST TEST", 20)
}
