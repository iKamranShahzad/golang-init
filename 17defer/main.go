package main

import "fmt"

func main() {
	defer fmt.Println("Hello there!")
	defer fmt.Println("there!")
	fmt.Println("Defer in go lang!")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}