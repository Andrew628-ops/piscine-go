package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	fromBase := len(baseFrom)
	toBase := len(baseTo)

	fromMap := make(map[rune]int)
	for i, c := range baseFrom {
		fromMap[c] = i
	}

	n := 0
	for _, c := range nbr {
		n = n*fromBase + fromMap[c]
	}

	if n == 0 {
		return string(baseTo[0])
	}

	result := ""
	for n > 0 {
		remainder := n % toBase
		result = string(baseTo[remainder]) + result
		n /= toBase
	}

	return result
}
