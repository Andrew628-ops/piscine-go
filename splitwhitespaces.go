package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	word := ""

	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
