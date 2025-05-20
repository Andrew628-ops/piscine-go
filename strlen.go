package piscine

func StrLen(s string) int {
	return len([]rune(s)) // Convert string to rune slice to get character count
}
