package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . input.txt output.txt airport-lookup.csv")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

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
		fields := strings.Fields(line)
		compacted := strings.Join(fields, " ")
		writer.WriteString(compacted + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

/*
a formatted tool, transfer the info in a .csv file that can be friendly by general customers
Information need to be converted.
The main parts: Airport name & dates and times

Airport: From ICAO(4-letters) / IATA(3-letters) to airport name
e.g:
##AYBK / #BUA, and the expected output: Buka Airport / Buka Airport

Dates: D(2007-04-05T12:30−02:00). Dates must be displayed in the output as DD-Mmm-YYYY format. E.g. "05 Apr 2007"
12 Hour time: T12(2007-04-05T12:30−02:00). These must be displayed as "12:30PM (-02:00)".
24 Hour time: T24(2007-04-05T12:30−02:00). These must be displayed as "12:30 (-02:00)".
Zulu time: T24(2007-04-05T12:30−02:00). These must be displayed as "12:30 (+00:00)".

Error handling
If the incorrect number of arguments are provided, display the usage.
If the input does not exist, your program must display "Input not found".
If the airport lookup does not exist, your program must display "Airport lookup not found".
If the airport data in question is malformed within the airport lookup, your program must display "Airport lookup malformed". The data is considered malformed if any column in the lookup is missing or blank.
In any case of an error, the output file must not be created or overwritten.

The first part: Search airport

A general concept:
input.txt -> prettifier -> output.txt

So firstly:
input.txt -> output.txt

And then add one more step:
input.txt -> prettifier, prettifier receive the value from input.txt, and checking if there is any value matched
prettifier return a value
output.txt receive the value from prettifier


*/
