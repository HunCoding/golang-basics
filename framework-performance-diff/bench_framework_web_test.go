package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func BenchmarkFrameworksWeb(b *testing.B) {

	ginGonic := "http://localhost:8080/testGinGonic"
	gorillaMux := "http://localhost:8081/testGorillaMux"
	goChi := "http://localhost:8082/testGoChi"
	goFiber := "http://localhost:8083/testGoFiber"

	type bench struct {
		name           string
		paramName      string
		resultExpected int64
		endpoint       string
		body           []byte
	}

	obj := ObjectExample{
		ID:   "test",
		Name: "TEST",
		Age:  100,
	}
	bt, _ := json.Marshal(obj)

	var tests []bench = []bench{
		{
			name:           "tests_with_the_right_value_GINGONIC",
			paramName:      "huncoding",
			resultExpected: 200,
			endpoint:       ginGonic,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_value_GINGONIC",
			paramName:      "TEST",
			resultExpected: 400,
			endpoint:       ginGonic,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_body_GINGONIC",
			paramName:      "huncoding",
			resultExpected: 400,
			endpoint:       ginGonic,
			body:           nil,
		},
		{
			name:           "tests_with_the_right_value_GORILLAMUX",
			paramName:      "huncoding",
			resultExpected: 200,
			endpoint:       gorillaMux,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_value_GORILLAMUX",
			paramName:      "TEST",
			resultExpected: 400,
			body:           bt,
			endpoint:       gorillaMux,
		},
		{
			name:           "tests_with_the_wrong_body_GORILLAMUX",
			paramName:      "huncoding",
			resultExpected: 400,
			endpoint:       gorillaMux,
			body:           nil,
		},
		{
			name:           "tests_with_the_right_value_GOCHI",
			paramName:      "huncoding",
			resultExpected: 200,
			endpoint:       goChi,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_value_GOCHI",
			paramName:      "TEST",
			resultExpected: 400,
			endpoint:       goChi,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_body_GOCHI",
			paramName:      "huncoding",
			resultExpected: 400,
			endpoint:       goChi,
			body:           nil,
		},
		{
			name:           "tests_with_the_right_value_GOFIBER",
			paramName:      "huncoding",
			resultExpected: 200,
			endpoint:       goFiber,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_value_GOFIBER",
			paramName:      "TEST",
			resultExpected: 400,
			endpoint:       goFiber,
			body:           bt,
		},
		{
			name:           "tests_with_the_wrong_body_GOFIBER",
			paramName:      "huncoding",
			resultExpected: 400,
			endpoint:       goFiber,
			body:           nil,
		},
	}

	go initGinGonic()
	go initGoChi()
	go initGoFiber()
	go initGorillaMux()

	time.Sleep(2 * time.Millisecond)

	for _, test := range tests {
		b.Run(test.name, func(bf *testing.B) {
			bf.ReportAllocs()
			for i := 0; i < bf.N; i++ {
				if result := callEndpoint(test.paramName, test.endpoint, test.body); result != test.resultExpected {
					bf.Errorf(`
					%s - Result is different than expected.
					expected=%d, 
					got=%d`,
						test.name,
						test.resultExpected,
						result,
					)
					return
				}
			}
		})
	}

}

func callEndpoint(param string, path string, body []byte) int64 {
	res, err := http.Post(
		fmt.Sprintf("%s/%s", path, param),
		"application/json",
		bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	return int64(res.StatusCode)
}
