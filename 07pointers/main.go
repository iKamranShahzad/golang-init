package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointors in goLang")
	
	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23

	var ptr = &myNumber // & is used to get the address of the variable
	fmt.Println("Value of myNumber is: ", ptr)
	fmt.Println("Value of myNumber is: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Value of myNumber is: ", myNumber)
}