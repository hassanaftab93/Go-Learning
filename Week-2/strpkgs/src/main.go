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
	//--

}
