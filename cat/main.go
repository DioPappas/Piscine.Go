package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			cat(os.Stdin)
		} else {
			printError("No input provided through pipe\n")
			os.Exit(1)
		}
		return
	}

	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			printError(err.Error() + "\n")
			continue
		}
		cat(file)
		file.Close()
	}
}

func cat(reader io.Reader) {
	content, err := io.ReadAll(reader)
	if err != nil {
		printError(err.Error() + "\n")
		return
	}
	for _, ch := range content {
		z01.PrintRune(rune(ch))
	}
}

func printError(msg string) {
	for _, ch := range msg {
		z01.PrintRune(ch)
	}
}
