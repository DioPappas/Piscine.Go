package main

import (
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(0)
	} else {
		for _, s := range args {
			file, err := os.Open(s)
			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(1)
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					os.Stderr.WriteString(err.Error() + "\n")
					os.Exit(1)
				} else {
					os.Stdout.Write(data)
				}
				file.Close()
			}
		}
	}
}
