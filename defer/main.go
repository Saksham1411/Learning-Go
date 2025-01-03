package main

import "fmt"

func main() {
	// defer is keyword which use to take the line of code to the end of the main function

	defer fmt.Print(" hello ")
	fmt.Print("worlddd ")
	// output worldd hello

	// if we have multiple defer then it works in Lifo format
	example()
	// hello 1 2 3 4
	// output worlddd 4 3 2 1 hello
}

func example(){
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}