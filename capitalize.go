package piscine

func Capitalize(s string) string {
	result := ""
	newWord := true

	for _, r := range s {
		if isAlphaNumeric(r) {
			if newWord && r >= 'a' && r <= 'z' {
				r = r - ('a' - 'A')
			} else if !newWord && r >= 'A' && r <= 'Z' {
				r = r + ('a' - 'A')
			}
			newWord = false
		} else {
			newWord = true
		}
		result += string(r)
	}

	return result
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
