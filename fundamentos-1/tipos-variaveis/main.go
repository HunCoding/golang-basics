package main

import "fmt"

func main() {

	var testInt8 int = 3
	var testInt16 int16 = 3
	var testInt32 int32 = 3
	var testInt64 int64 = 3
	var testString string = "test"
	var testFloat32 float32 = 32.3
	var testFloat64 float64 = 32.3
	var testinterface interface{} = 32.3

	fmt.Println(testInt8)
	fmt.Println(testInt16)
	fmt.Println(testInt32)
	fmt.Println(testInt64)
	fmt.Println(testString)
	fmt.Println(testFloat32)
	fmt.Println(testFloat64)
	fmt.Println(testinterface)
}
