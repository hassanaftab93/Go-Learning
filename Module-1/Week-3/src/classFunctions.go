package main

import (
	"fmt"

	"github.com/hassanaftab93/go-composite-data-types/pkgs/arrays"
	"github.com/hassanaftab93/go-composite-data-types/pkgs/hashTables"
	"github.com/hassanaftab93/go-composite-data-types/pkgs/maps"
	"github.com/hassanaftab93/go-composite-data-types/pkgs/slices"
	"github.com/hassanaftab93/go-composite-data-types/pkgs/structs"
	"github.com/hassanaftab93/go-composite-data-types/pkgs/variableSlices"
)

func function1() string {
	return "\nFunction call from classFunctions.go in package main\n"
}

func Init() {
	fmt.Println("\n", function1())
	fmt.Println(arrays.Called())
	fmt.Println(slices.Called())
	fmt.Println(variableSlices.Called())
	fmt.Println(hashTables.Called())
	fmt.Println(maps.Called())
	fmt.Println(structs.Called())
}

func Arrays() {
	fmt.Println("\nArray returned Nth Value: ", arrays.ReturnArray()[0]) //[N] at the end is the index of returned value
	fmt.Println("\nArray returned: ", arrays.ReturnArray()[:])           //[:] for whole array or leave out the [] part, for whole Array
}

func Slices() {
	fmt.Println("\nSlice1: arr[1:3]: ", slices.ReturnSlice1())
	fmt.Println("\nLen(Slice1): ", slices.ReturnLenSlice1())
	fmt.Println("\nCap(Slice1): ", slices.ReturnCapSlice1())
	fmt.Println("\nSlice2: arr[3:6]: ", slices.ReturnSlice2())
	fmt.Println("\nLen(Slice2): ", slices.ReturnLenSlice2())
	fmt.Println("\nCap(Slice2): ", slices.ReturnCapSlice2())
}

func VariableSlices() {
	slice := variableSlices.MakeIntSlice(5)
	fmt.Println("\nReturned Slice with Make(): ", slice)
	slice = variableSlices.AppendSlice(slice, 1)
	fmt.Println("\nReturned Slice with Append(): ", slice)
}

func Maps() {
	fmt.Println("\nMap: ", maps.ExportidMap())
	fmt.Println("\nMap value for \"Hassan\": ", maps.ExportidMap()["Hassan"])

	fmt.Println("\nAdding new pair: ")
	maps.AddKeyValuePair("TEO", 4)
	fmt.Println("\nMap after adding keyvalue pair: ", maps.ExportidMap())

	fmt.Println("\nDeleting a keyvalue pair")
	maps.DeleteKeyValuePair("idMap", "TEO")
	fmt.Println("\nMap after adding keyvalue pair: ", maps.ExportidMap())

	fmt.Println("\nIterating the Hash table:")
	maps.IterateTable()
}

func Structs() {
	fmt.Println("\n", structs.InitializePerson("Hassan", "Islamabad", "923135189591"))
	fmt.Println("\n", structs.InitializePerson("Joe", "NY", "001554782"))
}
