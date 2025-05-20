package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		// Check if the character is a digit
		if char < '0' || char > '9' {
			return 0 // Return 0 if a non-digit character is found
		}
		result = result*10 + int(char-'0')
	}
	return result
}
