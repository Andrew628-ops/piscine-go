package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0
	skip := false

	for _, char := range str {
		if char == ' ' {
			if count < 5 {
				continue // Skip spaces when building 5-char groups
			}
			// Keep space if it appears after a complete group
			result += string(char)
			continue
		}

		if skip {
			skip = false
			continue
		}

		result += string(char)
		count++

		if count == 5 {
			result += " "
			count = 0
			skip = true
		}
	}

	// Remove trailing space if exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
