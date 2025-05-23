package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 12 {
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		if result > (1<<31-1)/i {
			return 0
		}
		result *= i
	}

	return result
}
