package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("./newfile.txt")
	handleError(err)

	content := "Thor tufaan ka neta"
	length, err := io.WriteString(file, content)
	handleError(err)

	fmt.Println(length)
	defer file.Close()

	readFile("./newfile.txt")
}

func readFile(filename string){
	dataByte, err := os.ReadFile(filename)
	handleError(err)
	fmt.Println(string(dataByte))
	// data byte which is by default is a buffer need to convert it into string
}

func handleError(err error){
	if err!=nil {
		panic(err)
	}
}