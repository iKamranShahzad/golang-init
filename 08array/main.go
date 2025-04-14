package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in goLang!") // not really used in goLang, rather slices are used

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Cherry"

	fmt.Println("Value of fruits is: ", fruits)
	fmt.Println("Value of fruits is: ", len(fruits))

	var vegList = [5]string{"Potato", "" ,"Tomato", "Onion"}
	
	fmt.Println("Value of vegList is: ", vegList)
	fmt.Println("Value of vegList is: ", len(vegList)) // length is 5, doesn't depend on the number of elements rather the size of the array
}