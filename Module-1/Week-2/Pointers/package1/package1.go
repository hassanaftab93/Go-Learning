package package1

var Random int //Global variable for package1
var Random2 float32 = 3.0
var Random3 float32 = 6.0

func Export() int { //Function for returning a value
	return Random
}
