package video

import "os/exec"

// ffmpeg -i input.mp4 -vf "scale=iw*min(1280/iw\,720/ih):ih*min(1280/iw\,720/ih), pad=1280:720:(1280-iw*min(1280/iw\,720/ih))/2:(720-ih*min(1280/iw\,720/ih))/2" -c:a copy output.mp4
func Resize1280x720(id string) {
	output := id + "_resize.mp4"
	cmd := exec.Command("ffmpeg",
		"-i", id,
		"-vf", "scale=iw*min(1280/iw\\,720/ih):ih*min(1280/iw\\,720/ih), pad=1280:720:(1280-iw*min(1280/iw\\,720/ih))/2:(720-ih*min(1280/iw\\,720/ih))/2",
		"-c:a",
		"copy",
		"-y",
		output)
	cmd.CombinedOutput()

}
