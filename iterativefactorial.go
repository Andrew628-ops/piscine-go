package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= nb; i++ {
		if result > (1<<63-1)/i {
			return 0
		}
		result *= i
	}

	return result
}
