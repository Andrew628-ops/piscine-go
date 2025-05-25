package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fullPath := os.Args[0]

	lastSlash := -1
	for i, c := range fullPath {
		if c == '/' {
			lastSlash = i
		}
	}

	for i, c := range fullPath {
		if i > lastSlash {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
