package main

import (
	"fmt"
	"testing"
)

var num = 100

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

func BenchmarkPrimeNumbersCoding(b *testing.B) {
	// mock
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}

func BenchmarkPrimeNumbersImproved(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrimeImproved(v.input)
			}
		})
	}
}
