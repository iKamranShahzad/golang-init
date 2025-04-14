package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Hello, User!"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our app (between 1 and 5): ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n') // replace _ with err to see the error and vice versa
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T", input)
}