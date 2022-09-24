package main

import (
	"bufio" // for Input with spaces
	"fmt"
	"math"
	"os"      // for bufio
	"strings" // for strings
)

// Assignment 1
// Write a program which prompts the user to enter a floating point number
// and prints the integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.

func Assignment1() {

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

// Assignment 2
/*

Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive,
so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings,
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

*/

func Assignment2() {

	var temp string

	fmt.Print("\nEnter string: ")
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()

	line := userInput.Text()

	temp = strings.ToLower(line)

	if strings.HasPrefix(temp, "i") && strings.Contains(temp, "a") && strings.HasSuffix(temp, "n") {
		fmt.Println("\nFound!")
	} else {
		fmt.Println("\nNot Found!")
	}
}

func main() {

	Assignment1()
	Assignment2()
}
