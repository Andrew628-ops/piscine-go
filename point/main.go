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

	// Calculate ASCII values mathematically
	// 'x' = 120, ' ' = 32, '=' = 61, '4' = 52, '2' = 50,
	// ',' = 44, 'y' = 121, '1' = 49, '\n' = 10

	// Store values in the point struct itself
	points.x = 120 // Will be overwritten by setPoint, but we reuse the variable
	points.y = 32

	// Print sequence: x, space, =, space, 4, 2, comma, space, y, space, =, space, 2, 1, newline
	// Using the struct to store temporary values
	z01.PrintRune(rune(points.x)) // 'x' (120)
	z01.PrintRune(rune(points.y)) // ' ' (32)
	points.x = 61
	z01.PrintRune(rune(points.x)) // '=' (61)
	z01.PrintRune(rune(points.y)) // ' ' (32)
	points.x = 52
	z01.PrintRune(rune(points.x)) // '4' (52)
	points.x = 50
	z01.PrintRune(rune(points.x)) // '2' (50)
	points.x = 44
	z01.PrintRune(rune(points.x)) // ',' (44)
	z01.PrintRune(rune(points.y)) // ' ' (32)
	points.x = 121
	z01.PrintRune(rune(points.x)) // 'y' (121)
	z01.PrintRune(rune(points.y)) // ' ' (32)
	points.x = 61
	z01.PrintRune(rune(points.x)) // '=' (61)
	z01.PrintRune(rune(points.y)) // ' ' (32)
	points.x = 50
	z01.PrintRune(rune(points.x)) // '2' (50)
	points.x = 49
	z01.PrintRune(rune(points.x)) // '1' (49)
	points.x = 10
	z01.PrintRune(rune(points.x)) // '\n' (10)
}
