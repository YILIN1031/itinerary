package others

import (
	"fmt"
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
