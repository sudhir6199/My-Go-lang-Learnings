package main

import "fmt"

func main() {

	i := 0

	for { //Note Here no condition for infinite.

		fmt.Println(i)

		if i == 100 {
			break
		}

		i += 1

	}
}
