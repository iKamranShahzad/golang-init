package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in GOLANG!")
	fileContent := "This content is written by Kamran Shahzad"

	file, err := os.Create("./myTextFile.txt")

	checkNil(err)

	length, err := io.WriteString(file, fileContent)

	checkNil(err)

	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile()
}

func readFile() {
	databyte, err := os.ReadFile("./myTextFile.txt")
	checkNil(err)
	fmt.Println(string(databyte))
}

func checkNil(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}