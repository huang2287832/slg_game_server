
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go Proc(ctx, 1)
	go Proc(ctx, 2)
	go Proc(ctx, 3)
	go Proc(ctx, 4)

	time.Sleep(time.Second / 10)
	cancel()

	time.Sleep(time.Second)
}

func Proc(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Proc-%d ", n)
		}
	}
}
