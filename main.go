package main

import (
	"flag"
	"fmt"
	"os"

	"gitea.koodsisu.fi/yilinlai/itinerary/pkg/airport"
	"gitea.koodsisu.fi/yilinlai/itinerary/pkg/others"
)

func main() {
	argsCounter := len(os.Args) - 1
	help := flag.Bool("h", false, "show the help messsage")
	flag.Parse()

	if *help || argsCounter != 3 {
		others.UserHelper()
	}

	codeList := []string{"LAX", "JFK", "ABCD", "CDG"}
	//code := "ABCD"
	codeToName, err := airport.MappingCodeToAirportName("./airport-lookup.csv")
	if err != nil {
		fmt.Println("Error building airport map:", err)
		os.Exit(1)
	}

	for _, code := range codeList {
		if name, ok := codeToName[code]; ok {
			fmt.Printf("Code: %s, Name: %s.\n", code, name)
		} else {
			fmt.Println("No airport found for code:", code)
		}
	}

}

/*
1. 完成命令行界面，以及帮助说明
2. 读取.csv文件，完成数据转换
3.

*/
