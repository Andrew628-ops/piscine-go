package piscine

import "fmt"

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i])) // Print each character without a newline
	}
}
