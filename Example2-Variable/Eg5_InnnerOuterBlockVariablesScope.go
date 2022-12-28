//NOTE: outer block cannot access inner block variables where as inner block can access outer block variables.

package main
import ("fmt")

func main() {

	greet :="Hello"

	{
	   subject :="World"
	   serialnum:=1
	   fmt.Println("I said",greet,"to the",subject,serialnum)
	}
	//fmt.Println("I said",greet,"to the",subject,serialnum) //<---- This will throw error because subject & serialnum are inner block variables and outer block cannot access inner block variables where as inner block can access outer block variables.

}