package main

import (
	"context"
)

func main() {

	//Time1()

	// date := Time2()
	// log.Print("time is:", date)

	// ctx := context.Background()
	// err := Time3(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Success")
	// }

	//Time4()

	//Time5()

	k := "1592190385"
	ctx := context.WithValue(context.Background(), 0, k)
	x(ctx, k)
}
