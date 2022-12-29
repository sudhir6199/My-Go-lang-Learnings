package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ { //Note Here no braces or paranthesis around condition.

		if i == 3 {
			continue //Because of coninue it wont execute next print statement and skips to next iterations to print 4 and 5.
		}
		fmt.Println(i)
	}
}
