package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\nArray: ", arr)

	//index of 20 is 1
	x := 1

	//Lets try deleting 20 in this:
	slice1 := arr[:]
	slice2 := append(slice1[:x], slice1[x+1:]...)

	fmt.Println("\nSlice2 without 20 element:", slice2, "\n")
}
