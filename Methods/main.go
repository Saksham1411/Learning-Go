package main

import "fmt"

func main() {
	//no inhertiance in golang, NO super or parent
	// method and functions are same just diff is method is of a struct

	saksham := User{"saksham", "s@gmail.com", true, 21}
	fmt.Println(saksham)

	fmt.Printf("%+v\n", saksham)
	fmt.Printf("%+v\n", saksham.Name)

	saksham.GetName()

	saksham.SetName("sam")
	fmt.Printf("%+v\n", saksham)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetName (){
	fmt.Println("the name is", u.Name)
}

// passed the pointer here bcoz without pointer its passing as a copy 
// but to make actual change we have to pass by reference
func (u *User) SetName (newName string){
	u.Name = newName
	fmt.Println("new name is", u.Name)
}
