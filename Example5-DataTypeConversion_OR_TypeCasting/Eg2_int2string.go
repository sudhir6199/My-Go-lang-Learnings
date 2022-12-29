package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Integer to String
	value1 := 1
	fmt.Printf("\nvalue1 Value: %v Type is %T\n", value1, value1)

	value2 := strconv.Itoa(value1)
	fmt.Printf("\nvalue2 Value: %v Type is %T\n\n", value2, value2)
}
