package piscine

func Atoi(s string) int {
	// Handle empty string
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1

	// Handle potential signs: '+' or '-'
	if s[0] == '-' {
		sign = -1
		s = s[1:] // Remove the '-' sign
	} else if s[0] == '+' {
		s = s[1:] // Remove the '+' sign (no need to change sign)
	}

	for i := 0; i < len(s); i++ {
		// Only process digits
		if s[i] < '0' || s[i] > '9' {
			return 0 // Return 0 if non-digit character is found
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}
