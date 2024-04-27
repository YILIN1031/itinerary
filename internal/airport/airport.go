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

type AirportInfo struct {
	Icao string
	Iata string
	Name string
}

func CSVDataReader(filepath string) []AirportInfo {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil
	}
	headerMapping := make(map[string]int)
	for index, header := range headers {
		headerMapping[strings.TrimSpace(header)] = index
	}

	requiredColumns := []string{"icao_code", "iata_code", "name"}
	for _, column := range requiredColumns {
		if _, exists := headerMapping[column]; !exists {
			return nil
		}
	}

	var airports []AirportInfo

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
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

func isAirportDataCorrupted(name string) bool {
	onlyLettersRegex := regexp.MustCompile(`^[A-Za-z\s]+$`)
	if strings.TrimSpace(name) == "" {
		return true
	}
	for _, r := range name {
		if !unicode.IsPrint(r) || !onlyLettersRegex.MatchString(name) {
			return true
		}
	}
	return false
}

func AirportInfoPrettify(inputContent []byte, airportLookupFilepath string) (string, error) {
	var outputContent string
	var callbackErr error

	codeToName, err := MappingCodeToAirportName(airportLookupFilepath)
	if err != nil {
		others.UserHelper("4")
		return "", fmt.Errorf("missing data")
	}

	codeRegex := regexp.MustCompile(`#{1,2}[A-Z]{3,4}`)

	outputContent = codeRegex.ReplaceAllStringFunc(string(inputContent), func(code string) string {
		cleanCode := strings.TrimLeft(code, "#")

		name, ok := codeToName[cleanCode]
		if isAirportDataCorrupted(name) {
			callbackErr = fmt.Errorf("the code mapping to a corrupted name")
			others.UserHelper("4")
			return ""
		} else if ok {
			return name
		}
		return code
	})

	if callbackErr != nil {
		return "", callbackErr
	}

	return outputContent, nil
}
