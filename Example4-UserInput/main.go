package main

import (
	"fmt"
)

func main() {

// Single Input
	fmt.Print("\nEnter Your Name Go Developer: ")
	var devname string
	fmt.Scanf(%s",&devname)
	fmt.Print("\nHello ",devname)



// Two Inputs
	fmt.Printf("\n\n%s, for which firm do you work and how long? ",devname)
	var (
		company string
		experience int
	)
	fmt.Scanf("%s %d",&company,&experience)
	fmt.Print("\n",company,"!! That too ",experience,"! Its Awesomeee ",devname,"\n")


//Scanf return values
    fmt.Printf("\n\n%s, What is your age and location? ",devname)
	var (
		age int
		location string
	)
	count, error:=fmt.Scanf("%d %s",&age,&location)
	fmt.Print("\n\nCount:",count)
	fmt.Print("\nError:",error)
    fmt.Print("\n\n Thats interesting! ",devname)

}