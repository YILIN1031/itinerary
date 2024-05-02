package airport

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode"

	"gitea.koodsisu.fi/yilinlai/itinerary/internal/others"
)

// declare AirportInfo type
type AirportInfo struct {
	Icao string
	Iata string
	Name string
}

func CSVDataReader(filepath string) []AirportInfo {
	// open file firstly and close file after finishing all operations before return
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	// create a CSV reader and read the content of .csv file
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil
	}

	// create a mapping relationship called 'headerMapping',
	// mapping header string of the table to a index, e.g. name -> 0, icao_code -> 1 ...and so on
	headerMapping := make(map[string]int)
	for index, header := range headers {
		headerMapping[strings.TrimSpace(header)] = index
	}

	// check if the clolumn exists
	requiredColumns := []string{"icao_code", "iata_code", "name"}
	for _, column := range requiredColumns {
		if _, exists := headerMapping[column]; !exists {
			return nil
		}
	}

	// declare a variable of type AirportInfo to save final result
	var airports []AirportInfo

	for {
		// read one line at a time
		record, err := reader.Read()
		// quit when arrive the last line
		if err == io.EOF {
			break
		}
		// quit when error exists
		if err != nil {
			break
		}
		for _, column := range requiredColumns {
			// check if there is any empty cell when reading a line
			if strings.TrimSpace(record[headerMapping[column]]) == "" {
				return nil
			}
		}

		// assign the value to the property of 'airport' (airport.Icao, airport.Iata, airport.Name)
		airport := AirportInfo{
			Icao: record[headerMapping["icao_code"]],
			Iata: record[headerMapping["iata_code"]],
			Name: record[headerMapping["name"]],
		}
		airports = append(airports, airport)
	}
	return airports
}

func MappingCodeToAirportName(filepath string) (map[string]string, error) {
	airports := CSVDataReader(filepath)
	if airports == nil {
		return nil, fmt.Errorf("missing data")
	}

	// create a mapping relationship called 'codeMapping', mapping code to airport name,
	// e.g. HKG -> Hong Kong International Airport
	codeMapping := make(map[string]string)
	for _, airport := range airports {
		if airport.Iata != "" {
			codeMapping[airport.Iata] = airport.Name
		}
		if airport.Icao != "" {
			codeMapping[airport.Icao] = airport.Name
		}
	}
	return codeMapping, nil
}

// check if the airport name is normal unicode
func isAirportDataCorrupted(name string) bool {
	//onlyLettersRegex := regexp.MustCompile(`^[A-Za-z\s]+$`)
	for _, r := range name {
		//if !unicode.IsPrint(r) || !onlyLettersRegex.MatchString(name) {
		if !unicode.IsPrint(r) {
			return true
		}
	}
	return false
}

func AirportInfoPrettify(inputContent string, airportLookupFilepath string) (string, error) {
	var outputContent string

	codeToName, err := MappingCodeToAirportName(airportLookupFilepath)
	if err != nil {
		others.UserHelper("4")
		return "", err
	}

	// match the patterns of IATA or ICAO, start with "#" or "##",
	// e.g. IATA: #HKG, ICAO: ##VHHH
	codeRegex := regexp.MustCompile(`#{1,2}[A-Z]{3,4}`)

	// a simple example of closure function
	/*
		func ReplaceAllIntsFunc(ints []int, repl func(int) int) []int {
		    replacedInts := make([]int, len(ints))
		    for i, v := range ints {
		        replacedInts[i] = repl(v)
		    }
		    return replacedInts
		}

		func main() {
		    ints := []int{1, 2, 3, 4, 5}
		    replacedInts := ReplaceAllIntsFunc(ints, func(x int) int {
		        if x%2 == 0 {
		            return x * x
		        }
		        return x
		    })
		    fmt.Println(replacedInts)
		}
	*/
	// ReplaceAllStringFunc() works as a closure function
	outputContent = codeRegex.ReplaceAllStringFunc(inputContent, func(code string) string {
		cleanCode := strings.TrimLeft(code, "#")

		name, ok := codeToName[cleanCode]
		if ok && !isAirportDataCorrupted(name) {
			return name
		}
		return code
	})

	return outputContent, nil
}
