package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

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

	// Load the airport code to name mapping
	codeToName, err := airport.MappingCodeToAirportName(airportLookupFilepath)
	if err != nil {
		fmt.Println("Error building airport map:", err)
		return
	}

	// Read the entire input file
	inputContent, err := os.ReadFile(inputFilepath)
	if err != nil {
		fmt.Println("Fail to read the input file:", err)
		return
	}

	// Regular expression to find airport codes
	codeRegex := regexp.MustCompile(`\b([A-Z]{3,4})\b`) // Assuming all airport codes are 3 uppercase letters

	// Replace each airport code with its full name
	outputContent := codeRegex.ReplaceAllStringFunc(string(inputContent), func(code string) string {
		if name, ok := codeToName[code]; ok {
			return name
		}
		return code // return the original code if not found
	})

	// Write the modified content to the output file
	if err := os.WriteFile(outputFilepath, []byte(outputContent), 0644); err != nil {
		fmt.Println("Fail to write to the output file:", err)
		return
	}

	fmt.Println("Successful to map codes to airport names and write to output.txt")

}

/*
1. 完成基本命令行界面及帮助说明
2. 读取.csv文件，完成数据转换
3. 文件数据交换, 即将input.txt文件中的内容，传递到output.txt
4.
*/
