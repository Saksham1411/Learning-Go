package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter something")
	// comma, ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("you have entered this thing --> ", input)
}