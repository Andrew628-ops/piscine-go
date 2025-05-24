package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, c := range s {
		index := indexInBase(c, base)
		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' || seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func indexInBase(c rune, base string) int {
	for i, b := range base {
		if b == c {
			return i
		}
	}
	return -1
}
