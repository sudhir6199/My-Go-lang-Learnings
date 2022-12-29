package main

import "fmt"

func main() {

	i := 0
	for i < 5 { //Note Here no braces or paranthesis around condition.

		if i == 3 {
			break
		}
		fmt.Println(i)
		i += 1
	}
}
