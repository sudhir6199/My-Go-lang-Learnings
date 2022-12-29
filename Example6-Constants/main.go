package main

import (
	"fmt"
)

func main() {

	present := "vzm"
	const native = "vizag" //constant cannot be declared in shorthand like the var variable "present" declared above.
	//const native string = "vizag"      <----- same like var for constants also mentioning datatype is optional.

	fmt.Println(present)
	fmt.Println(native)

	present = "hyd"
	fmt.Println(present)

	/* Constant Variables cannot be updated, below one will throw error as variable "native" is a  constant.

	native = "vzm"
	fmt.Println(native)

	*/
}
