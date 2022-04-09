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

type TestStruct struct {
	Code        int    `json:"code"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Age         int    `json:"age"`
	Adress      Adress `json:"adress"`
	PhoneNumber string `json:"phone_number"`
}

type Adress struct {
	Street  int    `json:"street"`
	Number  int    `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
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
