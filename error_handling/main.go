package main

import (
	"fmt"
)

type ErrorTest struct {
	Code    int
	Message string
}

func (e ErrorTest) Error() string {
	return fmt.Sprintf("Erro ao processar dados, code=%d, message=%s",
		e.Code,
		e.Message,
	)
}

func main() {

	if err := test(); err != nil {
		fmt.Printf("%#v", err)
	}

}

func test() *ErrorTest {
	err := &ErrorTest{
		Code:    500,
		Message: "Deu erro aqui",
	}

	return err
}
