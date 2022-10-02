package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop.

Before entering the loop, the program should create an empty integer slice of size (length) 3.

During each pass through the loop,
the program prompts the user to enter an integer to be added to the slice.

The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order.

The slice must grow in size to accommodate any number of integers which the user decides to enter.

The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of
an integer.
*/

func Sort(input []int) []int { //										Sorting the slice
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	return input
}

func main() { //												Main - Get input from user and store in SORTED slice
	fmt.Println("\nAssignment 1 for Week 3")

	var input string
	var convInput int
	var slice = make([]int, 3)

	//															Everytime value is entered, it should be appended to slice.
	for iteration := 0; iteration < iteration+1; iteration++ {
		fmt.Println("\nEnter an integer:\tEnter \"x\" or \"X\" to exit")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			break
		} else {
			convInput, _ = strconv.Atoi(input)
			slice = append(slice, convInput)
			slice = Sort(slice) //										Sorts the slice in ascending order
			fmt.Println(slice)
		}
	}

}
