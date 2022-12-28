package main
import ("fmt")

func main() {
	var greet string = "Hello"
	var subject string = "World."
    var serial int = 1

	//Using Print
	fmt.Println("\nI said",greet,"to the",subject)

	fmt.Printf("\nI said %s to the %s number %s",greet,subject,serial) //Here %s is string format so doesnt support int i.e "serial" variable hold integer value not string, check the output!

	fmt.Printf("\nI said %v to the %v number %v",greet,subject,serial) //Here %v is variable default format so support the data type of variable passed as arugment i.e. although "serial" variable hold integer it supports, check the output!

    fmt.Printf("\nI said %d to the %d number %d",greet,subject,serial) //Here %d is decimal format so doesnt support string, check the output!

	fmt.Printf("\nFirst Var Type: %T Second Var Type: %T Third Var Type: %T",greet,subject,serial) //Here %T tells the datatype of the variable.
}