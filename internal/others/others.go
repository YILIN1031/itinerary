package others

import (
	"fmt"
)

func UserHelper(option string) {
	switch option {
	case "1":
		fmt.Printf("wrong number of command line arguments.\n")
		fmt.Printf("try 'go run . -h'.\n")
	case "2":
		fmt.Printf("input file does not exist.\n")
		fmt.Printf("try 'go run . -h'.\n")
	case "3":
		fmt.Printf("airport lookup not found.\n")
		fmt.Printf("try 'go run . -h'.\n")
	case "4":
		fmt.Printf("airport lookup malformed.\n")
	default:
		fmt.Printf("itinerary usage:\n")
		fmt.Printf("go run . ./input.txt ./output.txt ./airport-lookup.csv\n")
	}
}
