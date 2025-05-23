package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for candidates := nb; ; candidates++ {
		if IsPrime(candidates) {
		}
	}
}
