package timedate

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func FormatDate(originalDateStr string) string {
	dateStr := replaceZWithOffset(originalDateStr)
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		return originalDateStr
	}
	return t.Format("02 Jan 2006")
}

func FormatTime12(originalDateStr string) string {
	dateStr := replaceZWithOffset(originalDateStr)
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		return originalDateStr
	}
	_, offset := t.Zone()
	return fmt.Sprintf("%s (%+03d:%02d)", t.Format("03:04PM"), offset/3600, abs((offset%3600)/60))
}

func FormatTime24(originalDateStr string) string {
	dateStr := replaceZWithOffset(originalDateStr)
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		return originalDateStr
	}
	_, offset := t.Zone()
	return fmt.Sprintf("%s (%+03d:%02d)", t.Format("15:04"), offset/3600, abs((offset%3600)/60))
}

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
	timeDateRegex := regexp.MustCompile(`(D|T(12|24))\((.*?)\)`)
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
