package main

import (
	"errors"
	"fmt"
)

var Func1Err = errors.New("error from Func1")
var Func2Err = errors.New("error from Func2")
var Func3Err = errors.New("error from Func3")

func testeError() (int, error) {

	_, err := Func1()
	if err != nil {
		switch {
		case errors.Is(err, Func1Err):
			fmt.Println("error from expected function called")
		default:
			fmt.Printf("unexpected error: %s\n", err)
		}
	}

	_, err = Func2()
	if err != nil {
		switch {
		case errors.Is(err, Func2Err):
			fmt.Println("error from expected function called")
		default:
			fmt.Printf("unexpected error: %s\n", err)
		}
	}

	_, err = Func3()
	if err != nil {
		switch {
		case errors.Is(err, Func3Err):
			fmt.Println("error from expected function called")
		default:
			fmt.Printf("unexpected error: %s\n", err)
		}
	}

	return 1, nil
}

func Func1() (int, error) {
	return 1, errors.New("error from Func1")
}

func Func2() (int, error) {
	return 1, errors.New("error from Func2")
}

func Func3() (int, error) {
	return 1, errors.New("error from Func3")
}
