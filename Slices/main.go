package main

import (
	"fmt"
	"sort"
)

func main() {
	//slices are just like array but more usefull have diff inbuilt method

	var fruits = []string{"apple", "banana"}
	fmt.Println(fruits)

	fruits = append(fruits, "mango", "litchi")
	fmt.Println(fruits)

	fruits = append(fruits[1:]) //use to slice the slices
	//it will start the slice from 1st index uptill infinite
	fmt.Println(fruits)

	fruits = append(fruits[:2])
	// now this will starts from the 0th index to before 2nd index (2nd index is not inclusive)
	fmt.Println(fruits)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 944
	highScores[2] = 334
	highScores[3] = 634
	// highScores[4] = 634 // this will give error index out of bound

	highScores = append(highScores, 55, 777, 555) 
	//but this fine bcoz this realocate the memory
	
	sort.Ints(highScores) //use to sort arrays

	fmt.Println(highScores)
}
