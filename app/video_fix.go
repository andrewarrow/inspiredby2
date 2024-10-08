package app

import (
	"fmt"
	"os/exec"
)

// ffmpeg -i input.mp4 -c copy -f segment -segment_time 10 -reset_timestamps 1 output_%06d.mp4
func ProcessVideoFix2() {
	for i := 0; i < 335; i++ {
		output := fmt.Sprintf("data2/%06d.mov", i)
		cmd := exec.Command("ffmpeg",
			"-i", fmt.Sprintf("data2/%06d.mp4", i),
			"-y",
			output)
		fmt.Println(i)
		cmd.CombinedOutput()
	}
}

func ProcessVideoFix() {
	guid := "cd0bc6a1-a7aa-0b7d-d318-601f22783be8"

	minutes := 55
	//for i := 10; i < minutes+1; i++ {
	for i := 0; i < minutes; i++ {
		from := 0 + (i * 60)
		to := from + 10
		for j := 0; j < 6; j++ {
			//sectionId := fmt.Sprintf("%d_%d_%d", one["id"], i, j)
			output := fmt.Sprintf("data2/%s_%d_%d.mp4", guid, i, j)

			cmd := exec.Command("ffmpeg",
				"-ss", fmt.Sprintf("%d", from),
				"-i", "data2/"+guid+".mp4",
				"-to", fmt.Sprintf("%d", to),
				"-c:v", "libx264", "-c:a", "aac", "-y",
				//"-c", "copy",

				output)
			fmt.Println(i, j, from, to)
			cmd.CombinedOutput()

			from += 10
			to += 10
		}
	}

}
