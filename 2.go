package main

import (
	"log"
	"time"
)

func Time2() time.Time {
	millisec := time.Now().UnixNano()
	log.Print(millisec)
	return time.Unix(0, millisec)
}
