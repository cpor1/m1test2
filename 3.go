package main

import (
	"context"
	"fmt"
	"time"
)

func Time3(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("TIME OUT")
			cancel()
			return ctx.Err()
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("Running: ", i)
		}
	}
	fmt.Println("ALL DONE")
	return nil
}
