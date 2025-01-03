package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv package used to convert string into other DataTypes
	// strings.TrimSpace to remove \n jo ki enter press krne pe as input ajata

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("increasing +1 --> ", num+1)
	}

}
