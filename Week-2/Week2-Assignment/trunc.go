package main

import (
	"fmt"
	"math"
)

// Write a program which prompts the user to enter a floating point number
// and prints the integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.

func main() {

	var userNum float32
	var err error

	fmt.Print("\nEnter a Float value: ")
	_, err = fmt.Scan(&userNum)

	if err == nil {
		fmt.Println("\nTruncated Integer: ", math.Trunc(float64(userNum)))
	} else {
		fmt.Println("\nError: ", err)
	}
}
