package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	solution := make([]int, 0, 8)
	solve(solution)
}

func solve(solution []int) {
	if len(solution) == 8 {
		printSolution(solution)
		return
	}
	for pos := 1; pos <= 8; pos++ {
		if isSafe(solution, pos) {
			solve(append(solution, pos))
		}
	}
}

func isSafe(solution []int, pos int) bool {
	col := len(solution)
	for i := 0; i < col; i++ {
		if solution[i] == pos ||
			solution[i]-i == pos-col ||
			solution[i]+i == pos+col {
			return false
		}
	}
	return true
}

func printSolution(solution []int) {
	for _, digit := range solution {
		z01.PrintRune(rune(digit + '0'))
	}
	z01.PrintRune('\n')
}
