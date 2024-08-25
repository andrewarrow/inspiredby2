package browser

import (
	"fmt"
	"strconv"
)

func FormatSeconds(s string) string {
	seconds, _ := strconv.ParseFloat(s, 64)
	sec := int(seconds)

	hours := sec / 3600
	minutes := (sec % 3600) / 60
	newSeconds := sec % 60

	timecode := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, int(newSeconds))
	return timecode
}
