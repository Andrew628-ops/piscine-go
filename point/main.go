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

	z01.PrintRune(120) // 'x'
	z01.PrintRune(32)  // ' '
	z01.PrintRune(61)  // '='
	z01.PrintRune(32)  // ' '
	z01.PrintRune(52)  // '4'
	z01.PrintRune(50)  // '2'
	z01.PrintRune(44)  // ','
	z01.PrintRune(32)  // ' '
	z01.PrintRune(121) // 'y'
	z01.PrintRune(32)  // ' '
	z01.PrintRune(61)  // '='
	z01.PrintRune(32)  // ' '
	z01.PrintRune(50)  // '2'
	z01.PrintRune(49)  // '1'
	z01.PrintRune(10)  // '\n'
}
