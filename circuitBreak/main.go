package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/sony/gobreaker"
)

func mockService() (string, error) {
	if rand.Intn(100) > 90 {
		return "success", nil
	}

	return "", errors.New("error trying to process request")
}

func logState(name string, from gobreaker.State, to gobreaker.State) {
	fmt.Printf("%s: State change: %s -> %s\n", name, stateToString(from), stateToString(to))
}

func stateToString(state gobreaker.State) string {
	switch state {
	case gobreaker.StateClosed:
		return "Closed"
	case gobreaker.StateHalfOpen:
		return "Half-Open"
	case gobreaker.StateOpen:
		return "Open"
	default:
		return "Unknown"
	}
}

func main() {
	settings := gobreaker.Settings{
		Name:        "MyCircuitBreakerService",
		MaxRequests: 3,
		Interval:    5 * time.Second,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 2
		},
		OnStateChange: logState,
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	for i := 0; i < 10; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			return mockService()
		})

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Circuit Breaker: Success")
		}

		time.Sleep(1 * time.Second)
	}
}
