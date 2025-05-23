package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	for i := 2; i <= nb; i++ {
		// Check for potential overflow by ensuring the result won't exceed max int
		if result > (1<<31-1)/i {
			return 0
		}
		result *= i
	}

	return result
}
