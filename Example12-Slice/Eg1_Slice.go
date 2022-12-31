package main

import "fmt"

func main() {

	fmt.Print("\n\n")

	//Example 1
	arr := [12]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120}
	slice := arr[2:8] /* Here similar to Python it will read including 2nd index but do not print 8th index value so total 6 values in this example considered as SLICE LENGTH.
	Whereas SLICE CAPACITY is 10 because starting from 2nd index till end of list of the list its 10 elements.
	*/

	fmt.Println("Eg1:", arr)
	fmt.Println("Eg1:", slice)

	fmt.Println("Eg1:", len(slice))
	fmt.Println("Eg1:", cap(slice)) // Capacity is including the 8th index.

	fmt.Print("\n\n\n\n")

	//Example2: Try to understand below a bit tricky one!
	arr2 := [10]int{10, 20}
	slice2 := arr2[2:8]

	fmt.Println("Eg2:", arr2)
	fmt.Println("Eg2:", slice2)

	fmt.Println("Eg2:", len(slice2))
	fmt.Println("Eg2:", cap(slice2)) // Capacity is including the 8th index.

	fmt.Print("\n\n")

	//Example3: Initialze a slice
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("Eg3:", slice3)

	//Example4: Declare a slice
	slice4 := make([]int, 4, 7)
	fmt.Println("\nEg4:", slice4)
	slice4 = arr[:]
	fmt.Println("Eg4:", slice4, len(slice4), cap(slice4), "\n")

}
