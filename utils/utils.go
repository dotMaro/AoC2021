package utils

import (
	"io"
	"os"
	"strings"
)

// InputFile wraps os.Open.
// Panics on error.
func InputFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

// InputString returns the input file's content as a string.
func InputString(path string) string {
	file := InputFile(path)
	defer file.Close()
	inputBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(inputBytes)
}

// SplitInput returns the input file's contents split into lines.
func SplitInput(path string) []string {
	input := InputString(path)
	return SplitLine(input)
}

// SplitLine splits s into a slice of lines.
// It supports both with carriage return and without.
func SplitLine(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}
