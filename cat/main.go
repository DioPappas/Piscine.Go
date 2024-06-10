package main

import (
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("ERROR: Usage: go run . <filename>\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.Write(content)
}
