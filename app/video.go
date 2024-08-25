package app

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:11 -vframes 1 frame_2.jpg
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:00 -to 00:00:09 -c:v libx264 -c:a aac output.mp4
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 0 -to 9 -c:v libx264 -c:a aac output.mp4

func ProcessVideo(c *router.Context, guid string) {
	d, _ := getVideoDuration("data/" + guid + ".mp4")
	one := c.One("link", "where guid=$1", guid)
	c.FreeFormUpdate("update links set duration=$1 where guid=$2", d, guid)
	c.FreeFormUpdate("update links set photos_ready=true where guid=$1", guid)

	sec := int(d)
	minutes := (sec % 3600) / 60
	for i := 0; i < minutes+1; i++ {
		from := 0 + (i * 60)
		to := from + 10
		for j := 0; i < 6; j++ {
			cmd := exec.Command("ffmpeg", "-i", "data/"+guid+".mp4",
				"-ss", fmt.Sprintf("%d", from), "-to",
				fmt.Sprintf("%d", to),
				"-c:v", "libx264", "-c:a", "aac", "-y",
				fmt.Sprintf("data/%s_%d_%d.mp4", guid, i, j))
			fmt.Println(i, j, from, to)
			cmd.CombinedOutput()
			c.Params = map[string]any{}
			c.Params["link_id"] = one["id"]
			c.Params["section"] = fmt.Sprintf("%d_%d_%d", one["id"], i, j)
			c.Params["minute"] = i
			c.Params["sub"] = j
			c.Insert("link_section")
			//fmt.Println(string(b), err)
			from += 10
			to += 10
		}
	}

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
