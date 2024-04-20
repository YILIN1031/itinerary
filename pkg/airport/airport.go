package airport

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type AirportInfo struct {
	Icao string
	Iata string
	Name string
}

/*
CSV Data Reading Functionality:
Input file path, Returns a slice of the AirportInfo structure, facilitating subsequent data lookup,
i.e. CSVDataReader(filepath string, []AirportInfo).

Implementation Overview:
traverse the data in the .csv file
read the data from specific columns in each row
place the retrieved values into the AirportInfo structure
and then insert the AirportInfo type elements into the slice
*/

func CSVDataReader(filepath string) []AirportInfo {
	// Open the .csv file firstly
	file, err := os.Open(filepath)
	// If there are some exceptions when opening file, give error info and exit
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	// Waiting all the code behaviors and file operations finished and then closing the file before return
	defer file.Close()

	// Read the first line of .csv file at first
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return nil
	}
	// Creating a map relationshop between header text to header index
	headerMap := make(map[string]int)
	for index, header := range headers {
		headerMap[strings.TrimSpace(header)] = index
	}

	var airports []AirportInfo

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		airport := AirportInfo{
			Icao: record[headerMap["icao_code"]],
			Iata: record[headerMap["iata_code"]],
			Name: record[headerMap["name"]],
		}
		airports = append(airports, airport)
	}
	return airports
}

/*
Mapping ICAO / IATA codes to airport names
*/
func MappingCodeToAirportName(filepath string) (map[string]string, error) {
	airports := CSVDataReader(filepath)
	if airports == nil {
		return nil, fmt.Errorf("failed to read airport data")
	}

	codeMap := make(map[string]string)
	for _, airport := range airports {
		if airport.Iata != "" {
			codeMap[airport.Iata] = airport.Name
		}
		if airport.Icao != "" {
			codeMap[airport.Icao] = airport.Name
		}
	}
	return codeMap, nil
}
