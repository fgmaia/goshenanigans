package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go doStuff(ctx)

	fmt.Println("main thread stuff 1")
	time.Sleep(5000 * time.Millisecond)
	cancel()

	fmt.Println("main thread stuff 2")
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("main thread stuff 3")
	time.Sleep(5000 * time.Millisecond)
}

func doStuff(ctx context.Context) {
	for {
		if ctx.Err() == context.Canceled {
			fmt.Println("doing stuff got canceled")
			break
		}
		fmt.Println("doing stuff")
		time.Sleep(1000 * time.Millisecond)
	}
}
