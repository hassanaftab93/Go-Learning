package functions

import (
	"fmt"
)

func ForLoop(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("Loop count: %v \n", i+1)
	}
}
