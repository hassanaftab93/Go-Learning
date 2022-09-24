package main

import (
	"fmt"
	"math"
)

func main() {

	var userNum float32
	var err error

	fmt.Print("\nEnter a Float value: ")
	fmt.Scan(&userNum)

	if err == nil {
		fmt.Println("\nTruncated Integer: ", math.Trunc(float64(userNum)))
	} else {
		fmt.Println("\nError: ", err)
	}

}
