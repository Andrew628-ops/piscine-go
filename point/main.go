package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(p *point) {
	p.x = 6 * 7 // 42
	p.y = 3 * 7 // 21
}

func main() {
	var p point
	setPoint(&p)

	output := []rune{
		120, // 'x'
		32,  // ' '
		61,  // '='
		32,  // ' '
		52,  // '4'
		50,  // '2'
		44,  // ','
		32,  // ' '
		121, // 'y'
		32,  // ' '
		61,  // '='
		32,  // ' '
		50,  // '2'
		49,  // '1'
		10,  // '\n'
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
