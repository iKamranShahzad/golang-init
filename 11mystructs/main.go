package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang!")
	// no inheritance, no supuer, no parent, no classes

	kamran := User{"Kamran", "r4wdrake1@gmail.com", true, 25}
	fmt.Println("User Kamran: ", kamran)
	fmt.Printf("Kamran's details are: %+v\n", kamran)
	fmt.Printf("Name is %v and email is %v.", kamran.Name, kamran.Email)

}

type User struct {  // struct is a collection of fields, and the capital first letter makes it public as well as the fields
	Name string
	Email string
	Status bool
	Age int
}