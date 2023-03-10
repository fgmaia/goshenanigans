package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	go doStuff(ctx)

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("cancelling")
		cancel()
	case <-ctx.Done():
		fmt.Println("done")
		fmt.Println(ctx.Err())
	}

	for {
		fmt.Println("main tread loop")
		time.Sleep(2000 * time.Millisecond)
	}
}

func doStuff(ctx context.Context) {
	for {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("context.DeadlineExceeded")
			break
		} else if ctx.Err() == context.Canceled {
			fmt.Println("context.Canceled")
			break
		}
		fmt.Println("doing stuff")
		time.Sleep(1000 * time.Millisecond)
	}
}
