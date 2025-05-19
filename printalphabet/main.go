package main

import "github.com/01-edu/z01"

func main() {
	// loop through the ASCII values of 'a' to 'z'
	for r := 'a'; r <= 'z'; r++ {
		// print each character using z01.PrintRune
		z01.PrintRune(r)
	}
	// ptionally add a newline after printing all characters
	z01.PrintRune('\n')
}
