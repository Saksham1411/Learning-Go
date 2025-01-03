package main

import "fmt"

func main() {
	//no inhertiance in golang, NO super or parent

	saksham := User{"saksham", "s@gmail.com", true, 21}
	fmt.Println(saksham)

	fmt.Printf("%+v\n", saksham)
	fmt.Printf("%+v\n", saksham.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
