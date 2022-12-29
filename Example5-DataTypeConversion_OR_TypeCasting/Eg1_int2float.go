package main

import (
	"fmt"
)

func main() {

	var value1 int
	value1 = 1

	// Float to Integer
	var value2 float32 = float32(value1)

	fmt.Printf("\nvalue1 Value: %v Type is %T\n", value1, value1)
	fmt.Printf("\nvalue2 Value: %v Type is %T\n\n", value2, value2)

	// Integer to Float
	value3 := int(value2) //Here var value3 int = int(value2)
	fmt.Printf("\nvalue3 from value2 Value: %v Type is %T\n\n", value3, value3)
}
