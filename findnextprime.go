package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for candidates := nb; ; candidates++ {
		if isPrime(candidates) {
			return candidates
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
