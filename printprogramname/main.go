package main

import (
	"os"
	"path/filepath"

	"github.com/01-edu/z01"
)

func main() {
	progName := filepath.Base(os.Args[0])

	for _, c := range progName {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
