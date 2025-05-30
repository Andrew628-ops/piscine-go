package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printString("File name missing\n")
		return
	}

	if len(args) > 1 {
		printString("Too many arguments\n")
		return
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		printString("Error reading file\n")
		return
	}

	printString(string(content))
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
