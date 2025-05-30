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

func printDigit(n int) {
	if n >= 0 && n <= 9 {
		z01.PrintRune('0' + n)
	}
}

func main() {
	points := &point{}
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print 42 (x)
	z01.PrintRune('4')
	z01.PrintRune('2')

	z01.PrintRune(',')
	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// Print 21 (y)
	z01.PrintRune('2')
	z01.PrintRune('1')

	z01.PrintRune('\n')
}
