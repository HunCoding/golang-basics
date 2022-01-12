package main

import "fmt"

func main() {

	var testValue string = "Otavio"
	copyStringVALUE(testValue)
	fmt.Println(testValue)

	originalStringValue(&testValue)
	fmt.Println(testValue)

}

func copyStringVALUE(stringValue string) {
	stringValue = "TEST"
	fmt.Println(stringValue)
}

func originalStringValue(stringValue *string) {
	*stringValue = "TEST"
	fmt.Println(*stringValue)

}
