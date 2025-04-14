package main

import "fmt"

func main() {
	fmt.Println("Functions in go lang!")
	greet()
	// fmt.Println(proSum(10, 2, 3, 4, 5))

	mySum, myMessage := proSum(10, 20, 100, 12, 31)
	fmt.Println("Sum is: ", mySum)
	fmt.Println("Message is: ", myMessage)
}

func greet() {
	fmt.Println("Hello there!")
}

func proSum(values ...int) (int, string) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum, "Did you like it?"
}

