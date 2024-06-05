package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, res := range args {
		fmt.Println(res)
	}
}
