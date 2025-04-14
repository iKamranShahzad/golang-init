package main

import "fmt"

func main() {
	fmt.Println("Loops in go lang!")

	days := []string{"Sunday", "Monday", "Tueday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("Days of the week: ", days)

	for d:=0; d<len(days); d++ { // for loop in go
		fmt.Printf("Day %d of week is %v\n", d, days[d])
	}

	for index, day := range days { // index is the index of the array and day is the value at that index, used a lot
		fmt.Printf("Index is %v and value is %v\n",index, day)
	}

	rogueVal := 1

	for rogueVal < 10 {

		if rogueVal == 2 {
			goto lco
		}

		if rogueVal == 5 {
			break
		}

		fmt.Println("Rogue value is: ", rogueVal)
		rogueVal++
	}

	lco: // label for goto
	fmt.Println("I am inside lco label")

}