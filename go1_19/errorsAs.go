package main

import (
	"errors"
	"fmt"
)

func errorsAsError() {
	var newErr = errors.New("sentinel")
	var err = errors.New("foo")

	if errors.As(err, newErr) {
		fmt.Println("error here!")
	}
}
