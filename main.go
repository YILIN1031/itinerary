package main

import (
	"flag"
	"fmt"
	"os"

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

	/*
		outputContent, err := airport.AirportInfoPrettify(trimWhitespace, airportLookupFilepath)
		if err != nil {
			return
		}
	*/

	outputContent := timedate.TimeDatePrettify(trimWhitespace)

	if err := os.WriteFile(outputFilepath, []byte(outputContent), 0644); err != nil {
		fmt.Println("Fail to write to the output file:", err)
		return
	}

	fmt.Println("Successful")

	/*
		fmt.Printf("Testing for time date codes\n")
		dateTime := "2024-04-28T21:18-02:00"
		fmt.Println("Date:", timedate.FormatDate(dateTime))
		fmt.Println("12-hour Time:", timedate.FormatTime12(dateTime))
		fmt.Println("24-hour Time:", timedate.FormatTime24(dateTime))
	*/
}
