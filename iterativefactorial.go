package piscine

func IterativeFactorial(nb int) int {
	// Handle invalid inputs (negative numbers)
	if nb < 0 {
		return 0
	}

	// Factorial of 0 is 1
	if nb == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= nb; i++ {
		// Check for overflow before multiplying
		if result > (1<<63-1)/i {
			return 0
		}
		result *= i
	}

	return result
}
