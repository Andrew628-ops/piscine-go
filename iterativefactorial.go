package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		// Check for overflow before multiplying
		if result > 0 && (result > (1<<31-1)/i) {
			return 0 // overflow detected
		}
		result *= i
	}
	return result
}
