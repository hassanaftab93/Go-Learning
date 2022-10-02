package arrays

// Some variables
// var c [10]int
// var c2 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //Array Literal
// OR
// x:= [...]int{1,2,3,4...,n} 					//The [...] infers length from the "n" elements in array
// var x = [...]int{1, 2, 3}

func Called() string {
	return "arrays package called"
}

func ReturnArray() []int {
	var array [5]int
	for item := 0; item < len(array); item++ {
		array[item] = item + 5
	}
	return array[:]
}
