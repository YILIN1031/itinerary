package main

import (
	"flag"
	"fmt"
	"os"

	"gitea.koodsisu.fi/yilinlai/itinerary/internal/airport"
	"gitea.koodsisu.fi/yilinlai/itinerary/internal/others"
)

func main() {
	argsCounter := len(os.Args) - 1
	help := flag.Bool("h", false, "show the help messsage")
	flag.Parse()

	if *help {
		others.UserHelper("")
		return
	}

	if argsCounter != 3 {
		others.UserHelper("1")
		return
	}

	inputFilepath := os.Args[1]
	outputFilepath := os.Args[2]
	airportLookupFilepath := os.Args[3]

	// Read the entire input file
	inputContent, err := os.ReadFile(inputFilepath)
	if err != nil {
		others.UserHelper("2")
		return
	}

	airportLookupFile, err := os.Open(airportLookupFilepath)
	if err != nil {
		others.UserHelper("3")
		return
	}
	defer airportLookupFile.Close()

	outputContent, err := airport.AirportInfoPrettify(inputContent, airportLookupFilepath)
	if err != nil {
		return
	}

	// Write the modified content to the output file
	if err := os.WriteFile(outputFilepath, []byte(outputContent), 0644); err != nil {
		fmt.Println("Fail to write to the output file:", err)
		return
	}

	fmt.Println("Successful to map codes to airport names and write to output.txt")

}
