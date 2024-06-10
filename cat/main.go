package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		// If no filenames are provided, simply exit without error.
		return
	}

	for i := 1; i < len(os.Args); i++ {
		filename := os.Args[i]
		file, err := os.Open(filename)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		defer file.Close()

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
	}
}
