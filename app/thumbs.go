package app

import (
	"fmt"
	"os/exec"

	"github.com/andrewarrow/feedback/router"
)

// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:11 -vframes 1 frame_2.jpg
func extractFrameAt(index, seconds int, guid string) {
	output := fmt.Sprintf("%s_%d.jpg", guid, index)
	cmd := exec.Command("ffmpeg", "-i", "data/"+guid+".mp4",
		"-ss", fmt.Sprintf("%d", seconds),
		"-vframes", "1",
		"-y",
		output)
	cmd.CombinedOutput()

	output50 := fmt.Sprintf("%s_%d_50percent.jpg", guid, index)
	cmd = exec.Command("magic", output, "-resize",
		"50%",
		output50)
	cmd.CombinedOutput()
}
func ProcessThumbs(c *router.Context, guid string) {
	extractFrameAt(1, 1, guid)
	extractFrameAt(2, 11, guid)
}
