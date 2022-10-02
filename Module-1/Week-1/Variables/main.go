package main

import "fmt"

func main() {

	// Declare variable | If no value assigned then assigning type is a must

	var a int
	a = 5
	fmt.Printf("\nValue of 'a': %v\nType: %T\n", a, a)

	// Declare variable | Value assigned then assigning type is optional

	var b = 6
	fmt.Printf("\nValue of 'b': %v\nType: %T\n", b, b)

	// Declare variable | Value assigned then assigning type is optional

	var c = true
	fmt.Printf("\nValue of 'c': %v\nType: %T\n", c, c)

	// Declare variable | Value assigned then assigning type is optional

	var d bool //var d bool = false
	d = false
	fmt.Printf("\nValue of 'd': %v\nType: %T\n", d, d)

	// Declare variable | Value assigned then assigning type is optional

	var e = 6.2
	fmt.Printf("\nValue of 'e': %v\nType: %T\n", e, e)

	// Declare variable | Value assigned then assigning type is optional

	var f string = "f is a String"
	fmt.Printf("\nValue of 'f': %v\nType: %T\n", f, f)
}
