package piscine

func Index(s string, toFind string) int {
	lenS := len(s)
	lenFind := len(toFind)

	if lenFind == 0 {
		return 0
	}

	if lenFind > lenS {
		return -1
	}

	for i := 0; i <= lenS-lenFind; i++ {
		if s[i:i+lenFind] == toFind {
			return i
		}
	}
	return -1
}
