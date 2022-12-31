package main

import "fmt"

func main() {

	var list1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("\nlist1: ", list1, "\n", list1[2], "\n")

	list2 := [5]int{1, 2, 3}                           //Easy way of declaring
	fmt.Println("list2:", list2, "\n", list1[2], "\n") // This is printing 0 in the output list as list lenght is 5 but only 3 indexes filled with values.

	list3 := [5]string{"apple", "ball", "cat"}
	fmt.Println("list3:", list3, "\n", list3[2], list3[4], "\n") // This is printing empty space in the output list as list lenght is 5 but only 3 indexes filled with values.

}
