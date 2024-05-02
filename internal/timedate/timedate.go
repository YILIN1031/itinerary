package timedate

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// format input to normal date format
func FormatDate(originalDateStr string) string {
	dateStr := replaceZWithOffset(originalDateStr)
	// time.Parse: The first parameter is layout, the second one is source string
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		return originalDateStr
	}
	return t.Format("02 Jan 2006")
}

// format input to 12 hour time format
func FormatTime12(originalDateStr string) string {
	dateStr := replaceZWithOffset(originalDateStr)
	// time.Parse: The first parameter is layout, the second one is source string
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		return originalDateStr
	}
	_, offset := t.Zone()
	return fmt.Sprintf("%s (%+03d:%02d)", t.Format("03:04PM"), offset/3600, abs((offset%3600)/60))
}

// format input to 24 hour time format
func FormatTime24(originalDateStr string) string {
	dateStr := replaceZWithOffset(originalDateStr)
	// time.Parse: The first parameter is layout, the second one is source string
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		return originalDateStr
	}
	_, offset := t.Zone()
	return fmt.Sprintf("%s (%+03d:%02d)", t.Format("15:04"), offset/3600, abs((offset%3600)/60))
}

// check if a string ends with "Z"
func replaceZWithOffset(originalDateStr string) string {
	if strings.HasSuffix(originalDateStr, "Z") {
		originalDateStr = strings.TrimSuffix(originalDateStr, "Z") + "+00:00"
	}
	return originalDateStr
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TimeDatePrettify(inputContent string) string {
	var outputContent string
	// .*?
	// dot: a wildcard character that matches any single character (except newline).
	// asterisk: match the previous element (in this case any character) zero or more times
	// question mark, when it follows '*' or '+': It changes the matching mode to non-greedy mode.
	// This means that the expression will match as few characters as possible
	// instead of the default of matching as many characters as possible (i.e. greedy mode).
	// submatches[0]: (D|T(12|24))\((.*?)\)
	// submatches[1]: (D|T(12|24))
	// submatches[2]: (12|24)
	// submatches[3]: (.*?)
	timeDateRegex := regexp.MustCompile(`(D|T(12|24))\((.*?)\)`)

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
	outputContent = timeDateRegex.ReplaceAllStringFunc(inputContent, func(match string) string {
		submatches := timeDateRegex.FindStringSubmatch(match)
		prefix := submatches[1]
		timeStr := submatches[3]

		switch prefix {
		case "D":
			return FormatDate(timeStr)
		case "T12":
			return FormatTime12(timeStr)
		case "T24":
			return FormatTime24(timeStr)
		}
		return match
	})
	return outputContent
}
