package others

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

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

func WhitespacePrettify(inputContent string) string {
	convertedInputContent := strings.ReplaceAll(inputContent, "\\v", "\n")
	convertedInputContent = strings.ReplaceAll(convertedInputContent, "\\f", "\n")
	convertedInputContent = strings.ReplaceAll(convertedInputContent, "\\r", "\n")
	convertedInputContent = strings.ReplaceAll(convertedInputContent, "\\n", "\n")

	var result strings.Builder
	scanner := bufio.NewScanner(bytes.NewReader([]byte(convertedInputContent)))
	blankLines := 0

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			blankLines++
			if blankLines > 1 {
				continue
			}
		} else {
			blankLines = 0
		}
		fields := strings.Fields(trimmedLine)
		compacted := strings.Join(fields, " ")
		result.WriteString(compacted + "\n")
	}

	return result.String()
}
