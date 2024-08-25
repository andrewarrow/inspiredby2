package app

import (
	"fmt"
	"inspiredby2/google"
	"inspiredby2/groq"
	"os/exec"
	"strconv"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:11 -vframes 1 frame_2.jpg
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:00 -to 00:00:09 -c:v libx264 -c:a aac output.mp4
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 0 -to 9 -c:v libx264 -c:a aac output.mp4
// ffmpeg -i input_video.mp4 -vn -acodec pcm_s16le -ar 16000 -ac 1 output_audio.wav
// ffmpeg -i input_video.mp4 -b:a 32k -ar 16000 -acodec flac output_audio.flac

/*
rpc error: code = Unauthenticated desc = transport: per-RPC creds failed due to error: Post "https://oauth2.googleapis.com/token": read tcp [2606:8e80:2809:ef00:91fe:73a1:923f:7a03]:61562->[2607:f8b0:400a:800::200a]:443: read: no route to host
*/

func intCheck(a any, val int) bool {
	aa := fmt.Sprintf("%v", a)
	bb, _ := strconv.Atoi(aa)
	return bb == val
}

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
		for j := 0; j < 6; j++ {
			sectionId := fmt.Sprintf("%d_%d_%d", one["id"], i, j)
			oneSection := c.One("link_section", "where section=$1", sectionId)
			output := fmt.Sprintf("data/%s_%d_%d.mp4", guid, i, j)

			if len(oneSection) == 0 {

				cmd := exec.Command("ffmpeg", "-i", "data/"+guid+".mp4",
					"-ss", fmt.Sprintf("%d", from), "-to",
					fmt.Sprintf("%d", to),
					"-c:v", "libx264", "-c:a", "aac", "-y",
					output)
				fmt.Println(i, j, from, to)
				cmd.CombinedOutput()
				c.Params = map[string]any{}
				c.Params["link_id"] = one["id"]
				c.Params["section"] = sectionId
				c.Params["minute"] = i
				c.Params["sub"] = j
				c.Params["meta"] = 1
				c.Insert("link_section")
			}
			oneSection = c.One("link_section", "where section=$1", sectionId)
			//fmt.Println(string(b), err)

			flac := fmt.Sprintf("data/%s_%d_%d.flac", guid, i, j)
			if intCheck(oneSection["meta"], 1) {
				cmd := exec.Command("ffmpeg", "-i", output,
					"-b:a", "32k", "-ar", "16000", "-acodec", "flac",
					"-y",
					flac)
				fmt.Println(i, j, from, to)
				cmd.CombinedOutput()
				c.FreeFormUpdate("update link_sections set meta=2 where section=$1", sectionId)
			}
			oneSection = c.One("link_section", "where section=$1", sectionId)

			if intCheck(oneSection["meta"], 2) {
				stt := google.Speech(flac)
				c.FreeFormUpdate("update link_sections set meta=3,stt=$1 where section=$2",
					stt, sectionId)
			}

			// ---

			from += 10
			to += 10
		}
		minuteKey := fmt.Sprintf("%d_%d", one["id"], i)
		oneMinute := c.One("link_minute", "where minute_key=$1", minuteKey)

		if len(oneMinute) == 0 {

			all := c.All("link_section", "where link_id=$1 and minute=$2 order by sub", "",
				one["id"], i)
			buffer := []string{}
			for _, item := range all {
				buffer = append(buffer, item["stt"].(string))
			}
			s := groq.Summarize(strings.Join(buffer, " "))
			c.Params = map[string]any{}
			c.Params["link_id"] = one["id"]
			c.Params["minute"] = i
			c.Params["minute_key"] = minuteKey
			c.Params["summary"] = s
			c.Insert("link_minute")
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
