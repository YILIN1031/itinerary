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

	if *help || argsCounter != 3 {
		others.UserHelper()
	}

	inputFilepath := os.Args[1]
	outputFilepath := os.Args[2]
	airportLookupFilepath := os.Args[3]

	// Read the entire input file
	inputContent, err := os.ReadFile(inputFilepath)
	if err != nil {
		fmt.Println("Fail to read the input file:", err)
		return
	}

	outputContent := airport.AirportInfoPrettify(inputContent, airportLookupFilepath)

	// Write the modified content to the output file
	if err := os.WriteFile(outputFilepath, []byte(outputContent), 0644); err != nil {
		fmt.Println("Fail to write to the output file:", err)
		return
	}

	fmt.Println("Successful to map codes to airport names and write to output.txt")

}
