package main

import "fmt"

const LoginToken string = "asjdhfjhasdf" // Public constant can be accessed from other packages, declared with capital letter

func main() {
	// Type String
	var username string = "Kamran"
	fmt.Println("Hello, " + username + "!")
	fmt.Printf("Variable is of type: %T \n", username)

	// Type Boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// Type Integer (uint8)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// Type Float
	var smallFloat float64 = 255.455233241421
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default Values and some Aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit way of declaring variables
	var website = "learngolang.com"
	fmt.Println(website)

	// No var keyword way of declaring variables
	numberOfUsers := 3281  // only allowed inside functions
	fmt.Println("Number of users: ", numberOfUsers)

	// Login Token (const)
	fmt.Println("Login Token: ", LoginToken)
	fmt.Printf("Constant is of type: %T \n", LoginToken)
}