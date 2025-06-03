package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0
	skip := false
	charsProcessed := 0

	for _, char := range str {
		if char == ' ' {
			result += string(char)
			continue
		}

		charsProcessed++
		if skip {
			skip = false
			continue
		}

		if count == 5 {
			result += " "
			count = 0
			skip = true
			continue
		}

		result += string(char)
		count++
	}

	// Remove trailing space if exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
