package main

import "fmt"

func main() {

	list1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("\nlist1: ", list1)

	list1[0] = 9
	fmt.Println("\n Updated List: ", list1)
}
