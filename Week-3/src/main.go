package main

import (
	"fmt"

	"github.com/hassanaftab93/go-composite-data-types/pkgs/arrays"         //3.1.1
	"github.com/hassanaftab93/go-composite-data-types/pkgs/hashTables"     //3.1.1
	"github.com/hassanaftab93/go-composite-data-types/pkgs/maps"           //3.1.1
	"github.com/hassanaftab93/go-composite-data-types/pkgs/slices"         //3.1.2
	"github.com/hassanaftab93/go-composite-data-types/pkgs/structs"        //3.1.1
	"github.com/hassanaftab93/go-composite-data-types/pkgs/variableSlices" //3.1.3
)

func main() {
	fmt.Println("\nInitial commit to week 3")
	fmt.Println(arrays.Called())
	fmt.Println(slices.Called())
	fmt.Println(variableSlices.Called())
	fmt.Println(hashTables.Called())
	fmt.Println(maps.Called())
	fmt.Println(structs.Called())

	fmt.Println("\nCourse related calls:")
	//Class related funcs here
	fmt.Println(function1())

	fmt.Println("\nAssignment related calls:")
	//Assignment related calls here

}
