package slices

//Variables

//	A slice is underlying array of an array, A part of an Array. A Subset of an array
//	Can have variable sizes, however Arrays are variable
//	Every slice needs a Pointer, Length, Capacity

//For example

var arr = [...]string{"a", "b", "c", "d", "e", "f"}

func Called() string {
	return "slices package called"
}

func ReturnSlice1() []string {
	return arr[1:3]
}

func ReturnSlice2() []string {
	return arr[3:6]
}
