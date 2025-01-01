package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func performRequest(ctx context.Context, url string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if rand.Intn(3) != 0 {
			return errors.New("simulated request error")
		}
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		fmt.Println("Request succeeded with status:", resp.StatusCode)
		return nil
	}
}

func retryWithContextBackoff(ctx context.Context, url string, maxRetries int, baseDelay time.Duration, maxTimeout time.Duration) error {
	retryDelay := baseDelay

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := performRequest(ctx, url)
		if err == nil {
			fmt.Printf("Request succeeded on attempt: %d\n", attempt)
			return nil
		}

		jitter := time.Duration(rand.Int63n(int64(retryDelay)))
		sleepTime := retryDelay + jitter/2
		fmt.Printf("Attempt %d failed: %s. Retrying in %s...\n", attempt, err, sleepTime)

		select {
		case <-time.After(sleepTime):
			retryDelay *= 2
			if retryDelay > maxTimeout {
				retryDelay = maxTimeout
			}
		case <-ctx.Done():
			return fmt.Errorf("operation canceled: %w", ctx.Err())
		}
	}
	return fmt.Errorf("max retries reached for URL: %s", url)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	url := "https://jsonplaceholder.typicode.com/posts/1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := retryWithContextBackoff(ctx, url, 7, time.Millisecond*100, time.Second*2)
	if err != nil {
		fmt.Println("Operation failed:", err)
	}
}
