package main

import (
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 || args[0] != "-c" {
		os.Stderr.WriteString("Usage: go run . -c <number> <file> [file...]\n")
		os.Exit(1)
	}

	n, err := strconv.Atoi(args[1])
	if err != nil || n < 0 {
		os.Stderr.WriteString("Invalid byte count\n")
		os.Exit(1)
	}

	files := args[2:]
	if len(files) == 0 {
		os.Stderr.WriteString("No file specified\n")
		os.Exit(1)
	}

	errorOccurred := false

	for i, filename := range files {
		data, err := os.ReadFile(filename)
		if err != nil {
			os.Stderr.WriteString("open " + filename + ": " + err.Error() + "\n")
			errorOccurred = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				os.Stdout.WriteString("\n")
			}
			os.Stdout.WriteString("==> " + filename + " <==\n")
		}

		start := len(data) - n
		if start < 0 {
			start = 0
		}

		os.Stdout.Write(data[start:])
	}

	if errorOccurred {
		os.Exit(1)
	}
}
