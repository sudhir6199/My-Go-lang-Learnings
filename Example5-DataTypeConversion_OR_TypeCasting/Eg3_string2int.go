package main

import (
	"fmt"
	"strconv"
)

func main() {

	value2 := "123"

	//String to Integer
	value3, err := strconv.Atoi(value2) //value3 := strconv.Atoi(value2)
	fmt.Printf("\nvalue3 Value: %v Type is %T\n\n and ERRORs are: %v\n\n", value3, value3, err)

	value2 = "ab123"                    //is a string with both alphabates and numbers so this will fail
	value4, err := strconv.Atoi(value2) //value3 := strconv.Atoi(value2) is wrong syntax as Atoi always expects two variables i.e. err is second variable.
	fmt.Printf("\n\nvalue4 Value: %v Type is %T\n\n and ERRORs are: %v\n\n", value4, value4, err)
}
