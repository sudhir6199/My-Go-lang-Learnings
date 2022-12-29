package main

import "fmt"

func main() {

	var uinput int

	fmt.Printf("\nPlease enter a value for luck draw:")
	fmt.Scanf("%d", &uinput)

	if uinput == 9 {
		fmt.Printf("\nCongratulations you won a BIKE.\n\n")
	} else if uinput == 13 { //NOTE: Here else should be along with } go wont identify if mentioned in new line.
		fmt.Printf("\nCongratulations you won a TOY, try again for JACKPOT.\n\n")
	} else {
		fmt.Printf("\nOh no badluck, try again!\n\n")
	}

}
