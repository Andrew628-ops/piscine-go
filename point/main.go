package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	// Print the entire output in minimal PrintRune calls
	output := []rune("x = 42, y = 21\n")
	for _, c := range output {
		z01.PrintRune(c)
	}
}
