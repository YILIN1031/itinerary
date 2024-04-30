package main

import (
	"flag"
	"fmt"
	"os"

	"gitea.koodsisu.fi/yilinlai/itinerary/internal/airport"
	"gitea.koodsisu.fi/yilinlai/itinerary/internal/others"
	"gitea.koodsisu.fi/yilinlai/itinerary/internal/timedate"
)

func main() {
	help := flag.Bool("h", false, "show the help messsage")
	flag.Parse()

	if *help {
		others.UserHelper("")
		return
	}

	if flag.NArg() == 0 {
		others.UserHelper("0")
		return
	}

	if flag.NArg() != 3 {
		others.UserHelper("1")
		return
	}

	inputFilepath := flag.Arg(0)
	outputFilepath := flag.Arg(1)
	airportLookupFilepath := flag.Arg(2)

	inputContent, err := os.ReadFile(inputFilepath)
	if err != nil {
		others.UserHelper("2")
		return
	}

	trimWhitespace := others.WhitespacePrettify(inputContent)

	airportLookupFile, err := os.Open(airportLookupFilepath)
	if err != nil {
		others.UserHelper("3")
		return
	}
	defer airportLookupFile.Close()

	processAirportInfo, err := airport.AirportInfoPrettify(trimWhitespace, airportLookupFilepath)
	if err != nil {
		return
	}

	processDateTime := timedate.TimeDatePrettify(processAirportInfo)

	if err := os.WriteFile(outputFilepath, []byte(processDateTime), 0644); err != nil {
		fmt.Println("Fail to write to the output file:", err)
		return
	}

	fmt.Println("Successful")

}
