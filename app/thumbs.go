package app

import (
	"os/exec"

	"github.com/andrewarrow/feedback/router"
)

// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:11 -vframes 1 frame_2.jpg
func ProcessThumbs(c *router.Context, guid string) {
	output := guid + "_1.jpg"
	cmd := exec.Command("ffmpeg", "-i", "data/"+guid+".mp4",
		"-ss", "1",
		"-vframes", "1",
		"-y",
		output)
	cmd.CombinedOutput()
	output = guid + "_2.jpg"
	cmd = exec.Command("ffmpeg", "-i", "data/"+guid+".mp4",
		"-ss", "11",
		"-vframes", "1",
		"-y",
		output)
	cmd.CombinedOutput()
}
