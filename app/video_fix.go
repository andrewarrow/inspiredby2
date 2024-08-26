package app

import (
	"fmt"
	"os/exec"
)

func ProcessVideoFix() {
	guid := "cd0bc6a1-a7aa-0b7d-d318-601f22783be8"

	//ffmpeg -i input.mp4 -c copy -f segment -segment_time 10 -reset_timestamps 1 output_%06d.mp4
	//output_000009.mp4
	//

	//minutes := 55
	//for i := 10; i < minutes+1; i++ {
	for i := 10; i < 11; i++ {
		from := 0 + (i * 60)
		to := from + 10
		for j := 0; j < 6; j++ {
			//sectionId := fmt.Sprintf("%d_%d_%d", one["id"], i, j)
			output := fmt.Sprintf("data/%s_%d_%d.mp4", guid, i, j)

			cmd := exec.Command("ffmpeg",
				"-ss", fmt.Sprintf("%d", from),
				"-i", "data/"+guid+".mp4",
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
