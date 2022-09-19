package package1

import (
	"strconv"
)

var Second int = 6
var x int = 4

const num int = 2

var Imaginary complex64

var ExportedValue int = x << num  // Bitwise << Left Shift Operator 1 0 0 becomes 1 0 0 0 0
var ExportedValue2 int = x >> num // Bitwise >> Right Shift Operator 1 0 0 becomes 0 0 1

var Str string = "5"

func Calculate(Str string) int {
	z := 0
	z, _ = strconv.Atoi(Str)
	return z
}

func MakeComplex(x float32, y float32) complex64 {
	return complex(x, y)
}

func MakeComplex2() complex64 { //Using global package1 variable in package1.go
	return complex(Random2, Random3)
}
