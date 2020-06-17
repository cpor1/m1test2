package main

import (
	"log"
	"time"
)

func Time1() {
	i := 0
	for i < 3 {
		log.Print("time now: ", time.Now().UnixNano()/int64(time.Millisecond))
		time.Sleep(3 * time.Second)
		i++
	}
}
