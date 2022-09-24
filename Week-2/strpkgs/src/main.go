package main

import (
	"fmt"
	"strings-project/pkg/functions"
	"strings-project/pkg/stringpackage"
)

func main() {
	//Initial code and function calling from 2 packages
	fmt.Println("\nThis is an output string")
	fmt.Println("\nFunctions package (add):", functions.Add(1, 2))
	fmt.Println("\nString package (concat):", stringpackage.Concat("Hello ", "Hassan"))

	//Testing the weird function
	fmt.Println("\nThe weird function called (1): ", functions.TestFunction(1))
	fmt.Println("\nThe weird function called (2): ", functions.TestFunction("?"))
	fmt.Println("\nThe weird function called (3): ", functions.TestFunction(1.2))
	//It worked, buet yet to figure out, how to return dynamic values from a function
}
