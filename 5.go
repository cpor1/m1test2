package main

import (
	"log"
	"strconv"
	"time"
)

func Time5() {
	i, _ := strconv.ParseInt("1592190294764144364", 10, 64)
	tm := time.Unix(i, 0)
	log.Print(tm.Weekday())
	log.Print(int(tm.Weekday()) + 1)
}
