package main

import (
	"fmt"
	"strconv"
	"time"
)

func Time4() {
	i, _ := strconv.ParseInt("1592190385", 10, 64)
	tm := time.Unix(i, 0)
	fmt.Println(tm.Minute())
}
