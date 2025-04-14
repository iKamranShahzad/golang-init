package main

import "fmt"

func main() {
	fmt.Println("Methods in go lang!")
	// no inheritance, no supuer, no parent, no classes

	kamran := User{"Kamran", "r4wdrake1@gmail.com", true, 25}
	fmt.Println("User Kamran: ", kamran)
	fmt.Printf("Kamran's details are: %+v\n", kamran)
	fmt.Printf("Name is %v and email is %v.", kamran.Name, kamran.Email)
	kamran.GetStatus()
	kamran.NewMail()
	fmt.Printf("Name is %v and email is %v.", kamran.Name, kamran.Email)
}

type User struct {  // struct is a collection of fields, and the capital first letter makes it public as well as the fields
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.Status)
}

func (u User) NewMail() {  // this is a method, not a function, because it is associated with a struct, and it only changes the copy of the value not actual value
	u.Email = "\nTestKamran@umail.com"
	fmt.Println("\nEmail has been changed to: ", u.Email)
}