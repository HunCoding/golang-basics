package main

import "fmt"

func main() {

	testMultipleReturnString, testMultipleReturnInt := returnTwoValues()
	testMultipleNamedValuesString, testMultipleNamedReturnInt := returnNamedValues()

	fmt.Println(testMultipleNamedReturnInt, testMultipleNamedValuesString)
	fmt.Println(testMultipleReturnInt, testMultipleReturnString)

}

func returnTwoValues() (string, int) {

	return "TEST", 20

}

func returnNamedValues() (returnStringValue string, returnIntValue int) {

	returnIntValue = 20
	returnStringValue = "TEST NAMED RETURN VALUE"

	return

}
