package functions

import (
	"reflect"
)

func Add(x, y int) int {
	return x + y
}

// Function for Returning dynamic variable types // DYNAMIC FUNCTION
// Figure out how to return dynamic values

func TestFunction(x interface{}) reflect.Type {
	switch g := x.(type) {
	case string:
		return reflect.TypeOf(g)
		//return g
	case int:
		return reflect.TypeOf(g)
		//return g
	default:
		return reflect.TypeOf(g)
		//return g
	}
}
