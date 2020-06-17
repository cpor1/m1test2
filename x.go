package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func x(ctx context.Context, k string) {
	for i := 0; i < 3; i++ {
		now := time.Now()
		ts, _ := strconv.ParseInt(k, 10, 64)
		timeFromTS := time.Unix(ts, 0)
		fmt.Println(now)
		fmt.Println(timeFromTS)
		diff := now.Sub(timeFromTS)
		fmt.Println("Diff between to date: ", diff)
		time.Sleep(3 * time.Second)
	}
}
