package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		"testKey",
		"testValue",
	)

	printUntilCancel(ctx)
}

func printUntilCancel(ctx context.Context) {
	fmt.Println(ctx.Value("testKey"))
}
