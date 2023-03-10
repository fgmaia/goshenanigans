package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	go doStuff(ctx)

	fmt.Println("main thread stuff 1")
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("main thread stuff 2")
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("main thread stuff 3")
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("main thread stuff 4")
	time.Sleep(5000 * time.Millisecond)
}

func doStuff(ctx context.Context) {
	for {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("doing stuff got deadline")
			break
		}
		fmt.Println("doing stuff")
		time.Sleep(1000 * time.Millisecond)
	}
}
