package main

import (
	"fmt"
	"strings"
)

func isPrintableASCII(r rune) bool {
	return (r >= 32 && r <= 126) || r != '\n'
}

func formatAsciiArt(input string, asciiArt map[rune][]string) (string, error) {
	if input == "" {
		return "", nil
	}

	// Replace literal "\n" with actual newline characters
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Check for non-printable ASCII characters
	for i, char := range input {
		if !isPrintableASCII(char) && char != '\n' {
			return "", fmt.Errorf("bad request: invalid character '%c' at position %d", char, i)
		}
	}

	var result strings.Builder
	lines := strings.Split(input, "\n")
	lastLineIdx := len(lines) - 1

	// Process each line of input
	for lineIdx, line := range lines {
		if line == "" {
			if lineIdx < lastLineIdx {
				result.WriteString("\n")
			}
			continue
		}

		// Process each of the 8 vertical lines of ASCII art
		for row := 0; row < 8; row++ {
			for _, char := range line {
				if art, exists := asciiArt[char]; exists {
					result.WriteString(art[row])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
