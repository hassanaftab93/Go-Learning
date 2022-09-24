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

	string1 := "Hassan is studying BSCS from Hassan University, Hassan Abdal"
	string2 := "  These are Extra Spaces.  "
	string3 := "??sampleEmail@email.com??"

	fmt.Println("\nFunction output: ", stringpackage.ReplaceIt(string1, "Hassan", "Kanza", 1))

	fmt.Println("\nTrim Function output: ")

	fmt.Println(stringpackage.SpaceTrimmer(string2))  //Trims Spaces
	fmt.Println(stringpackage.Trimmer(string3, "??")) //Trims Leading and Trailing characters

}
