package main

import (
	"os"
	"strings"
)

// loadAsciiArt loads ASCII art characters from a specified file and returns a map with each character's art.
func loadAsciiArt(filename string) (map[rune][]string, error) {
	// Read the file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Normalize all line endings consistently for all files
	normalizedContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	normalizedContent = strings.ReplaceAll(normalizedContent, "\r", "\n")

	// Split the content by newline
	lines := strings.Split(normalizedContent, "\n")
	// Initialize the map to store ASCII art for each character
	asciiArt := make(map[rune][]string)

	// Iterate through the lines in blocks to map each character to its ASCII art
	for i := 1; i < len(lines); i += 9 {
		// Ensure there are enough lines for the 8-line ASCII art block
		if i+8 >= len(lines) {
			break
		}
		// Determine the character represented by this block (ASCII starting from space, code 32)
		char := rune(32 + i/9)
		// Map the character to its ASCII art lines (each character has 8 lines of art)
		asciiArt[char] = lines[i : i+8]
	}

	// Return the map of characters and their ASCII art
	return asciiArt, nil
}
