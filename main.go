package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var input string

	fmt.Println("\nVerification Code: ")
	fmt.Scanln(&input)

	file, err := ioutil.ReadFile("code.txt")
	if err != nil {
		return
	}
	filecontents := string(file)

	if input == filecontents {
		fmt.Println("\nCode Is Correct")
	} else {
		fmt.Println("\nCode Is Incorrect")
		return
	}
}
