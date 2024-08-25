package app

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:11 -vframes 1 frame_2.jpg
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:00 -to 00:00:09 -c:v libx264 -c:a aac output.mp4

func ProcessVideo(c *router.Context, guid string) {
	d, _ := getVideoDuration("data/" + guid + ".mp4")
	c.FreeFormUpdate("update links set duration=$1 where guid=$2", d, guid)
}

func getVideoDuration(filePath string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	durationStr := strings.TrimSpace(string(output))
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
