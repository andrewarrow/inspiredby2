package video

import (
	"os/exec"
)

func RemoveBottom(id string) {
	//ffmpeg -i input.mp4 -vf "crop=in_w:in_h-110:0:0" -c:a copy output.mp4
	output := id + "_remove_bottom.mp4"
	cmd := exec.Command("ffmpeg",
		"-i", id,
		"-vf", "crop=in_w:in_h-51:0:0", "-c:a", "copy",
		"-y",
		output)
	cmd.CombinedOutput()

}
