package piscine

func LoafOfBread(str string) string {
	word := ""
	result := ""

	nonSpaceCount := 0
	skipping := false

	for _, r := range str {
		if r == ' ' {
			continue
		}

		if skipping {
			skipping = false
			continue // skip this character
		}

		word += string(r)
		nonSpaceCount++

		if nonSpaceCount == 5 {
			result += word + " "
			word = ""
			nonSpaceCount = 0
			skipping = true // set flag to skip next character
		}
	}

	if result == "" {
		return "Invalid Output\n"
	}

	return result[:len(result)-1] + "\n" // remove last space, add newline
}
