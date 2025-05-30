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

	message := "x = 42, y = 21\n"
	for _, r := range message {
		z01.PrintRune(r)
	}
}
