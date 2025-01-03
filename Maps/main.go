package main

import "fmt"

func main() {

	myMap := make(map[string]string)

	myMap["JS"] = "JAVAScript"
	myMap["cpp"] = "c plus plus"
	myMap["py"] = "python"

	fmt.Println(myMap)
	fmt.Println(myMap["JS"]) // geting a value for a particular key
	
	delete(myMap, "cpp") // deleting a key valur pair
	fmt.Println(myMap)

	// loops

	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}
}