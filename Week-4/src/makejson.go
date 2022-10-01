package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.

Your program should create a map and add the name and address to the map
using the keys “name” and “address”, respectively.

Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.


*/

func getName() string {
	fmt.Println("\nEnter Name: ")
	userinput := bufio.NewScanner(os.Stdin)
	userinput.Scan()
	output := userinput.Text()
	return output
}

func getAddr() string {
	fmt.Println("\nEnter Address: ")
	userinput := bufio.NewScanner(os.Stdin)
	userinput.Scan()
	output := userinput.Text()
	return output
}

func main() {
	fmt.Println("\nWeek 4 - Assignment 1")
	getName()
	getAddr()
}
