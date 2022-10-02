package main

import (
	"fmt"
	"project/package1" // Importing local package
	"reflect"          // Used for finding TypeOf
)

/*
	IMPORTANT:
		1. Inorder to call variables/function from local package,
			the Variables and Functions must start with a Capital Letter e.g RandomInteger, Export()
		2. `package NAME` should be same within one FOLDER PACKAGE
			'on compilaion, compiler makes a single package out of multiple .go files'
		3. %s to replace variables in Strings
*/

func main() {

	//Pointers
	var y = 5
	var randomSa *int = &y
	var i int = *randomSa

	//Assigning Values to Pointers
	variable := new(int)
	*variable = 9

	//Complex Numbers
	var comp complex64 = complex(2, 3)

	//Printing Values
	fmt.Println("\nVariable i -> y Value: ", i)
	fmt.Println("\nVariable variable -> Value: ", *variable)
	fmt.Println("\nComplex Number: ", comp) //Printing a complex number

	//Calling values from the local Package and Printing
	fmt.Println("\nImported Variable Value: ", package1.Random)
	fmt.Println("\nImported Variable Value from 2nd File in package1: ", package1.Second)
	fmt.Println("\nImported Function Output from package1.Export(): ", package1.Export())

	var Name string = "Hassan"
	fmt.Printf("\nHi, %s\n", Name) // Using %s

	fmt.Println("\nImported Value Output from package1: ", package1.ExportedValue)
	fmt.Println("\nImported Value Output from package1: ", package1.ExportedValue2)
	fmt.Println("\nType of Var from Calculate(\"5\"): ", reflect.TypeOf(package1.Calculate("5")))

	fmt.Println("\nComplex Func Output for MakeComplex(2.0,3.0) is:", package1.MakeComplex(2.0, 3.0))
	fmt.Println("\nComplex Func Output for MakeComplex2() is:", package1.MakeComplex2())
}
