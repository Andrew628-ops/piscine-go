package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -n {
			printMinInt()
			return
		}
		n = -n
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func printMinInt() {
	minIntStr := "-2147483648"

	for _, r := range minIntStr {
		z01.PrintRune(r)
	}
}
