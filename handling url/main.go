package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//this is like fetch or axios.get in js
	url := "https://dummyjson.com/c/3503-8b51-40d5-8d92"
	res, err := http.Get(url)
	handleError(err)

	defer res.Body.Close()
	dataByte, err := io.ReadAll(res.Body)
	handleError(err)
	data := string(dataByte)
	// data is need to convert into string
	fmt.Println("response", data)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
