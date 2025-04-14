package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Web App")
	fmt.Println("Please rate our app between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our app", input)

	// Conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //TrimSpace removes leading and trailing white spaces

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}