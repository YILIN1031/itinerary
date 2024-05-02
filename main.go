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
	// defines a command line flag 'h', it is a boolean flag with a default value of 'false'.
	// "show the help message" severs as the description of the flag
	help := flag.Bool("h", false, "show the help messsage")

	// parses the program's command line arguments
	// and updates the flag variables (in this case, help) based on the values provided.
	// After parsing, you can check the value of the flag
	// by dereferencing help (because flag.Bool returns a pointer to a boolean).
	// it is commonly used in command-line tools to allow users to request help information through the -h options.
	flag.Parse()

	// check if the user has used the flag,
	// dereference the bool type pointer to get the value
	if *help {
		others.UserHelper("")
		return
	}

	// check if the number of arguments is 0, namely doesn't add any argument in command line
	if flag.NArg() == 0 {
		others.UserHelper("0")
		return
	}

	// if the argument number in commend-line is not equal to 3 (after excluding then 'h' mark),
	// than quit the program
	if flag.NArg() != 3 {
		others.UserHelper("1")
		return
	}

	inputFilepath := flag.Arg(0)         // assign the first argument to input.txt filepath
	outputFilepath := flag.Arg(1)        // assign the second argument to output.txt filepath
	airportLookupFilepath := flag.Arg(2) // assign the third argument to airport-lookup.csv filepath

	if inputFilepath == outputFilepath || inputFilepath == airportLookupFilepath || outputFilepath == airportLookupFilepath {
		fmt.Println("Error: File paths must not be the same.")
		return
	}

	// open the input.txt file firstly,
	// input the filepath and get a return value of []byteA as inputContent
	inputContent, err := os.ReadFile(inputFilepath)
	if err != nil {
		others.UserHelper("2")
		return
	}

	// process the blank lines and whitespace of input.txt
	trimWhitespace := others.WhitespacePrettify(string(inputContent))

	// check weather airport-lookup.csv can be opened normally
	airportLookupFile, err := os.Open(airportLookupFilepath)
	if err != nil {
		others.UserHelper("3")
		return
	}
	// Note: os.ReadFile() vs os.Open()
	// When using os.Open(), you gain finer-grained control over the file,
	// allowing you to perform multiple read and write operations,
	// which is suitable for handling large files
	// or situations that require multiple reads and writes.
	// When using os.ReadFile(), you can read an entire file more succinctly,
	// which is suitable for quickly and once-only reading the entire file.
	// A simple and basic example about write operation of os.Open():
	/*
			file, err := os.Open("example.txt")
		    if err != nil {
		        fmt.Println("Error opening file:", err)
		        return
		    }
		    defer file.Close()

		    buf := make([]byte, 1024) // create a buffer of size 1024 bytes
		    for {
		        n, err := file.Read(buf)
		        if err == io.EOF {
		            break
		        }
		        if err != nil {
		            fmt.Println("Error reading file:", err)
		            return
		        }
		        fmt.Print(string(buf[:n]))
		    }
	*/

	// The defer keyword is used to schedule subsequent function calls to be executed
	// when the function returns.
	// Here, it is used to schedule airportLookupFile.Close() to be called
	// at the end of the current function,
	// ensuring that the opened file is properly closed.
	// This is a common practice to prevent resource leaks.
	defer airportLookupFile.Close()

	// to process airport data,
	// pass in the input.txt file with processed spaces
	// and the file path of airport-lookup.csv as parameters.
	processAirportInfo, err := airport.AirportInfoPrettify(trimWhitespace, airportLookupFilepath)
	if err != nil {
		return
	}

	// process the format of date and time
	processDateTime := timedate.TimeDatePrettify(processAirportInfo)

	// write all the things to output.txt
	// File Permission: 0644 -> 0-User-Group-Others -> 0-RWX-RWX-RWX -> 0-110-100-100
	// R: Read, W: Write, X: eXecute
	if err := os.WriteFile(outputFilepath, []byte(processDateTime), 0644); err != nil {
		fmt.Println("Fail to write to the output file:", err)
		return
	}

	//fmt.Println("Successful, please check your output.txt now.")

}
