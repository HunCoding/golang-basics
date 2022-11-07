package main

import "fmt"

func log() {
	fmt.Printf(
		"action=FindUserByIdAndName, message=init, userId=%s, userName=%s, error=Error trying to get user by id and name, trace=MongoDB timeout",
		"123455",
		"HUNCODING",
	)
}
