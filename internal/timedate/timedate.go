package timedate

import (
	"fmt"
	"strings"
	"time"
)

// 时间日期转换
// 输入: []byte, 输出string
// 标记： D(...), T12(...), T24(...)
// 日期：2022-05-09 -> 09 May 2022
// 时间：03:09Z / 17:54-11:00 -> 03:09 (+00:00) / 17:54 (-11:00)
// 例子：
// D(2022-05-09T20:00+12:00) -> 09 May 2022
// T12(2022-05-09T20:00+12:00) -> 08:00AM (+12:00)
// T24(2022-05-09T20:00+12:00) -> 20:00 (+12:00)

// 解析并格式化日期
func FormatDate(dateStr string) string {
	dateStr = replaceZWithOffset(dateStr)
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	return t.Format("02 Jan 2006")
}

// 解析并格式化12小时制时间
func FormatTime12(dateStr string) string {
	dateStr = replaceZWithOffset(dateStr)
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return ""
	}
	_, offset := t.Zone()
	return fmt.Sprintf("%s (%+03d:%02d)", t.Format("03:04PM"), offset/3600, abs((offset%3600)/60))
}

// 解析并格式化24小时制时间
func FormatTime24(dateStr string) string {
	dateStr = replaceZWithOffset(dateStr)
	t, err := time.Parse("2006-01-02T15:04-07:00", dateStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return ""
	}
	_, offset := t.Zone()
	return fmt.Sprintf("%s (%+03d:%02d)", t.Format("15:04"), offset/3600, abs((offset%3600)/60))
}

// 将字符串中的 "Z" 替换为 "+00:00"，以满足time.Parse的要求
func replaceZWithOffset(dateStr string) string {
	if strings.HasSuffix(dateStr, "Z") {
		dateStr = strings.TrimSuffix(dateStr, "Z") + "+00:00"
	}
	return dateStr
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
