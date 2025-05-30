package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Stdout.WriteString("File name missing\n")
		return
	}
	if len(args) > 1 {
		os.Stdout.WriteString("Too many arguments\n")
		return
	}

	data, err := os.Open(args[0])
	if err != nil {
		return
	}
	defer data.Close()

	io.Copy(os.Stdout, data)
}
