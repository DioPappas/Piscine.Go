package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := args[0]

	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Print(string(content))
}
