package main

// func main() {
// 	ctx, cancel := context.WithCancel(
// 		context.Background(),
// 	)
// 	go printUntilCancel(ctx)

// 	cancel()

// 	time.Sleep(10 * time.Second)
// }

// func printUntilCancel(ctx context.Context) {
// 	count := 0
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Cancel signal received, exiting")
// 			return
// 		default:
// 			time.Sleep(1 * time.Second)
// 			fmt.Printf("Printing until cancel, number = %d \n", count)
// 			count += 1
// 		}
// 	}
// }
