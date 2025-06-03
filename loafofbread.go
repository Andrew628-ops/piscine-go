package piscine

func LoafOfBread(str string) string {
	cleaned := ""
	for _, r := range str {
		if r != ' ' {
			cleaned += string(r)
		}
	}

	if len(cleaned) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(cleaned); {
		if len(cleaned)-i < 5 {
			break
		}
		result += cleaned[i:i+5] + " "
		i += 6 // Take 5 chars, then skip 1
		count++
	}

	if count == 0 {
		return "Invalid Output\n"
	}

	return result[:len(result)-1] + "\n" // remove last space and add newline
}
