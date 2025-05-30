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

	ascii := [15]int{
		120, // x
		32,  //
		61,  // =
		32,  //
		52,  // 4
		50,  // 2
		44,  // ,
		32,  //
		121, // y
		32,  //
		61,  // =
		32,  //
		50,  // 2
		49,  // 1
		10,  // \n
	}

	i := 0
	for i < 15 {
		z01.PrintRune(ascii[i])
		i++
	}
}
