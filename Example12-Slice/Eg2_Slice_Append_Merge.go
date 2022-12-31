package main

import "fmt"

func main() {
	arr := [12]int{10, 20, 30, 40, 50}
	slice := arr[2:8]

	fmt.Println("\n", arr)
	fmt.Println("\nslice:", slice)

	//Append slice
	slice2 := append(slice, 90, 80, 70, 60)
	fmt.Println("\nslice2: slice + elements =>", slice2, "\n")

	//Merge Slice
	slice3 := append(slice, slice2...)
	fmt.Println("\nslice3: slice2 + slice =>", slice3, "\n")
}
