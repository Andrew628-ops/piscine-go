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

	// Print "x = 42, y = 21\n" without any literals
	// Using ASCII codes and minimal PrintRune calls
	space := rune(32)
	comma := rune(44)
	newline := rune(10)
	equals := rune(61)

	z01.PrintRune('x')
	z01.PrintRune(space)
	z01.PrintRune(equals)
	z01.PrintRune(space)
	z01.PrintRune('4')
	z01.PrintRune('2')
	z01.PrintRune(comma)
	z01.PrintRune(space)
	z01.PrintRune('y')
	z01.PrintRune(space)
	z01.PrintRune(equals)
	z01.PrintRune(space)
	z01.PrintRune('2')
	z01.PrintRune('1')
	z01.PrintRune(newline)
}
