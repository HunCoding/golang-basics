package main

import (
	"errors"
	"fmt"
)

func main() {

	if test, err := funcaoTest(); err != nil {
		fmt.Println(test, err)
	}

	if test, err := funcaoTest2(); err != nil {
		fmt.Println(test, err)
	}
}

func funcaoTest() (string, error) {
	return "", errors.New("Test")
}

func funcaoTest2() (string, error) {
	return "test", nil
}
