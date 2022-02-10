package main

import "fmt"

type User[T any, B any] struct {
	name T
	age  B
}

func testComparable[T comparable](arg1 T, arg2 T) bool {
	return arg1 == arg2
}

type NumberTest interface {
	int64 | float64 | float32
}

func testComparingNumbers[T NumberTest](arg1 T, arg2 T) bool {
	return arg1 > arg2
}

type TestImplementInt int64

func testTokenImplement[T ~int64](arg1 T) {
	fmt.Println(arg1)
}

func main() {

	userTest := User[string, int64]{
		name: "test",
		age:  20,
	}
	fmt.Println(userTest)

	fmt.Println(testComparable(20, 20))

	var testNumberInt64 int64 = 20
	fmt.Println(testComparingNumbers(testNumberInt64, testNumberInt64))

	var testImplement TestImplementInt = 30
	testTokenImplement(testImplement)

}
