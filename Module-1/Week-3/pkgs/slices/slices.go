package slices

//Variables

//	A slice is underlying array of an array, A part of an Array. A Subset of an array
//	Can have variable sizes, however Arrays are variable
//	Every slice needs a Pointer, Length, Capacity

//For example

var arr = [...]string{"a", "b", "c", "d", "e", "f"}

func Called() []string {
	slice := []string{"Slice ", "Package ", "Called"} //This is a slice literal, with an empty [], compiler thinks this is a slice,however it starts at the start of array and ends with the end of array

	return slice
}

func ReturnSlice1() []string {
	return arr[1:3]
}

func ReturnLenSlice1() int {
	return len(arr[1:3])
}

func ReturnCapSlice1() int {
	return cap(arr[1:3])
}

func ReturnSlice2() []string {
	return arr[3:6]
}

func ReturnLenSlice2() int {
	return len(arr[3:6])
}

func ReturnCapSlice2() int {
	return cap(arr[3:6])
}
