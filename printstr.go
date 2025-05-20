package piscine

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		// Print each character, followed by a newline (for clarity if needed)
		print(string(s[i]))
	}
}
