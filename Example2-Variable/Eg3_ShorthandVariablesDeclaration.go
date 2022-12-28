//NOTE: Doesn't work for Constants

package main
import ("fmt")

func main() {


	//Way1: ShortHand Variables Initalization
	var greet1, subject1 string

	//Declaring Variables
	greet1 = "Hello"
	subject1 = "World1."

	fmt.Println("I said",greet1,"to the",subject1)


	//Way2: ShortHand Variables Initalization
	var greet2, subject2 string

	//Declaring Variables
	greet2,subject2 = "Hello","World2."

	fmt.Println("I said",greet2,"to the",subject2)	



	//Way3: ShortHand Variables Initalization
	var greet3, subject3 string = "Hello","World3."
	fmt.Println("I said",greet3,"to the",subject3)


	//Way4: ShortHand Variables Initalization
	var (
		greet4 string= "Hello" 
		subject4 string ="World" //No comma at end
		serialnum int =1
	)
	fmt.Println("I said",greet4,"to the",subject4, serialnum)

}