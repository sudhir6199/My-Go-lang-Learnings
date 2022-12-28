package main
import ("fmt")

//greet :="Hello World"  //<------ This shorthand can be used only within function and cannot be used for global variable declaration else will lead to error in go version 1.19.3 darwin/amd64. 
var greet string = "Hello" //This is a global variable.
//fmt.Println("I said",greet) //Global variables cannot be printed the Print statement works only in main function.

func main() {

	   subject :="World."  // This is a local variable.
	   fmt.Println("I said",greet,"to the",subject)
}
//fmt.Println("I said",greet,"to the",subject)    //<-----Uncommenting will throw error as "subject" variable is a local variable.