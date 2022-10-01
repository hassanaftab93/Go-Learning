package variableSlices

/*
Variable Slices - What are they?

- We can make slices/arrays using make() function

-- 2 Arg version: make(type,length) 			<- Here the length = capacity
--- e.g slice = make([]int, 10)

-- 3 Arg version: make(type,length,capacity) 	<- Here you specify length and capacity separately
--- e.g slice = make([]int,10,15)

- Add elements to end of slice? we use append()
-- Lets we have a case of a slice with Len: 0
--- slice = make([]int,0,3) 					<- Here the length is 0, and length of underlying array is 3, we can append an item into this slice which indirectly makes the array bigger
--- slice = append(slice,100)					<- appends the value "100" into the slice and underlying array
*/
func Called() string {
	return "variableSlices package called"
}

func MakeIntSlice(size int) []int {
	slice := make([]int, size)
	return slice
}

func AppendSlice(sl []int, element int) []int {
	sl = append(sl, element)
	return sl
}
