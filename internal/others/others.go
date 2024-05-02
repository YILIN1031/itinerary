package others

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// display help information based on different options
func UserHelper(option string) {
	switch option {
	case "0":
		fmt.Printf("enter 'go run . -h' to show more information.\n")
	case "1":
		fmt.Printf("wrong number of command line arguments.\n")
	case "2":
		fmt.Printf("input file does not exist.\n")
	case "3":
		fmt.Printf("airport lookup not found.\n")
	case "4":
		fmt.Printf("airport lookup malformed.\n")
	default:
		fmt.Printf("itinerary usage:\n")
		fmt.Printf("go run . <./input.txt> <./output.txt> <./airport-lookup.csv>\n")
	}
}

// process the blank lines and whitespace
func WhitespacePrettify(inputContent string) string {
	// "strings" contains functions for manipulating strings.
	// These functions include searching
	// and replacing strings, comparing, trimming, splitting,
	// and joining strings, etc.
	// For example, you can use strings.
	// Split to divide a string or strings.
	// Join to concatenate an array of strings.
	convertedInputContent := strings.ReplaceAll(inputContent, "\\v", "\n")
	convertedInputContent = strings.ReplaceAll(convertedInputContent, "\\f", "\n")
	convertedInputContent = strings.ReplaceAll(convertedInputContent, "\\r", "\n")
	convertedInputContent = strings.ReplaceAll(convertedInputContent, "\\n", "\n")

	// use strings.Builder to build a string that can be manipulated
	var result strings.Builder

	scanner := bufio.NewScanner(bytes.NewReader([]byte(convertedInputContent)))

	// set a blank line counter
	blankLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		// remove the leading and trailing spaces firstly
		trimmedLine := strings.TrimSpace(line)
		// if line is in blank, add 1 to counter
		if trimmedLine == "" {
			blankLines++
			if blankLines > 1 {
				continue
			}
		} else {
			blankLines = 0
		}
		// split and then combine
		fields := strings.Fields(trimmedLine)
		compacted := strings.Join(fields, " ")
		result.WriteString(compacted + "\n")
	}

	return result.String()
}
