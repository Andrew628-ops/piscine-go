package piscine

import (
	"github.com/01-edu/z01"
)

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	count := AlphaCount(s)

	for _, r := range []rune(IntToString(count)) {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func IntToString(n int) string {
	if n == 0 {
		return "0"
	}

	str := ""
	digits := []rune("0123456789")

	for n > 0 {
		remainder := n % 10
		str = string(digits[remainder]) + str
		n /= 10
	}
	return str
}
