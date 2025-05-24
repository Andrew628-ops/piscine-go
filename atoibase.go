package piscine

func indexInbase(r rune, base string) int {
	for i, br := range base {
		if br == r {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0
	for _, r := range s {
		digit := indexInbase(r, base)
		if digit == -1 {
			return 0
		}
		result = result*baseLen + digit
	}
	return result
}
