package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO!")

	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Pineapple")
	fmt.Println("Fruit List is: ", fruitList)

	fruitList = append(fruitList[0:2], "") // ending index is exclusive
	fmt.Println("Fruit List Now is: ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234 
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
//	highScores[4] = 321 // will throw an error because the size of slice is 4 (default allocated size)
	fmt.Println("High Scores: ", highScores)

	highScores = append(highScores, 555, 666, 777) // will increase the size of slice
	fmt.Println("High Scores Now: ", highScores)

	sort.Ints(highScores) // sort the slice
	fmt.Println("Sorted High Scores: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // check if the slice is sorted

	// Slicing a slice based on index

	var courses = []string{"react", "angular", "vue", "svelte", "ruby"}
	fmt.Println("Courses: ", courses)
	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println("Courses after removing index 2: ", courses)

}