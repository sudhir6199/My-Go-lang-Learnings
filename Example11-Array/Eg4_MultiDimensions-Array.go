package main

import "fmt"

func main() {

	list1 := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 0}} //2 ROWS x 5 COLUMNS

	fmt.Println("\nlist1: ", list1)

	fmt.Println("\nList Row1: ", list1[0])
	fmt.Println("List Row2: ", list1[1])

	fmt.Println("\nList Row1 Column1: ", list1[0][0])
	fmt.Println("List Row2 Column1: ", list1[1][0])
}
