package main

import (
	"fmt"
)

// Custom types used
type sliceIntArray [10]int

/* Code Principles
1. Understandability is Important
	1.1. If you are requested to find a feature, if found easily, then Code is understandable
	1.2. Where data is used/found, Traceability should be easy
	1.3. Maybe data is bad/wrong in a file being read, Issue should be traceable

2. Debugging Principles
	2.1. There are 2 Possibilities
		2.1.1. Function is written incorrectly (Logic needs to be revised, try Dry running)
		2.1.2. Maybe the Data being passed into the function is wrong. (Try tracing back, Console logs can help)

3. Supporting debugging
	3.1. Functions need to be understandable
	3.2. Is the function performing its purpose correctly?
	3.3. Data needs to be traceable
		3.3.1. Global Variable complicate this
		3.3.2. Try to code in a way such that all the data is going Func -> Func
*/

/* How to Name Functions/Vars/Consts
1. Function Naming
	- Understandable names
	- Parameter names
	- Functional Cohesion is important | one Function should perfrom one Operation.
*/

/*
	Assignment Task

Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program,
you should write a function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.

The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice.
You should write a Swap() function which performs this operation.

Your Swap() function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice.

The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/
func GetInputFromUser(sliceInt *sliceIntArray) {
	fmt.Println("\nEnter 10 Integer Values: ")
	var temp int
	for index := 0; index < len(*sliceInt); index++ {

		fmt.Print("Enter Value", index+1, ": ")
		fmt.Scan(&temp)
		(*sliceInt)[index] = temp
	}

	fmt.Println("\n10 Values Entered")
}

func BubbleSort(sliceInt *sliceIntArray) { //Takes a slice of Integeres
	for i := 0; i < len(*sliceInt)-1; i++ {
		for j := 0; j < len(*sliceInt)-i-1; j++ {
			Swap(&(*sliceInt)[j], &(*sliceInt)[j+1])
		}
	}
}

func Swap(sliceItem1, sliceItem2 *int) { //Swaps two adjacent values Args: 1. slice of ints 2. index i for position of slice
	if *sliceItem1 > *sliceItem2 {
		*sliceItem1, *sliceItem2 = *sliceItem2, *sliceItem1
	}
}

func main() {
	fmt.Println("\nModule-2 Week-1 BubbleSort Assignment")
	var slice sliceIntArray //Declaring a slice of size 10

	fmt.Print("\nLength of Slice: ", len(slice))

	GetInputFromUser(&slice)
	fmt.Println("\nBefore Sorting the Values: ", slice)

	BubbleSort(&slice)
	fmt.Println("\nAfter Sorting the Values: ", slice)
	//Display the Numbers in one line after Sorting

}
