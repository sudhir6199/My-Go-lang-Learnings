package main
import ("fmt")

func main() {

	//Way4: Dynamic Writing: ShortHand Variables Declaration without mentioning Datatype as it get added by compiler in this case.

	greet4 :="Hello"  //NOTE: this way of declaration works only within main() function cannot used to declare global variable i.e. define variable outside main() function.

	subject4 :="World"

	serialnum:=1
	serialnum=2
	//serialnum="hey"     //this is wrong as serialnum variable initial type is integer.

	fmt.Println("I said",greet4,"to the",subject4, serialnum)

}