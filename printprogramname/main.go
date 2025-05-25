package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	fmt.Println(programName)
}
