package main

import "fmt"

func main() {
	hello()

	res := add(3,4)
	fmt.Println(res)
	
	proRes := proAdd(1,2,3,5,6,7)
	fmt.Println(proRes)


}

func proAdd(values ...int) int {
	sum := 0
	for _, val := range(values){
		sum += val

	}
	return sum
}

func add(a int, b int) int {
	return a + b
}

func hello() {
	fmt.Println("helllooo")
}