package app

import (
	"fmt"
	"inspiredby2/google"
	"inspiredby2/video"
	"os"
	"os/exec"
	"strconv"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:11 -vframes 1 frame_2.jpg
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 00:00:00 -to 00:00:09 -c:v libx264 -c:a aac output.mp4
// ffmpeg -i cd0bc6a1-a7aa-0b7d-d318-601f22783be8.mp4 -ss 0 -to 9 -c:v libx264 -c:a aac output.mp4
// ffmpeg -i input_video.mp4 -vn -acodec pcm_s16le -ar 16000 -ac 1 output_audio.wav
// ffmpeg -i input_video.mp4 -b:a 32k -ar 16000 -acodec flac output_audio.flac

/*
rpc error: code = Unauthenticated desc = transport: per-RPC creds failed due to error: Post "https://oauth2.googleapis.com/token": read tcp [2606:8e80:2809:ef00:91fe:73a1:923f:7a03]:61562->[2607:f8b0:400a:800::200a]:443: read: no route to host
*/

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
func intCheck(a any, val int) bool {
	aa := fmt.Sprintf("%v", a)
	bb, _ := strconv.Atoi(aa)
	return bb == val
}

var BUCKET = "/Users/aa/bucket/"

func ProcessVideo(c *router.Context, guid string) {
	project := c.One("project", "where guid=$1", guid)
	file := project["file"].(string)
	d, _ := video.GetVideoDuration(BUCKET + file)
	sec := int(d)
	minutes := (sec % 3600) / 60
	for i := 0; i < minutes+1; i++ {
		from := 0 + (i * 60)
		to := from + 3
		for j := 0; j < 20; j++ {
			processSectionOfOrigVideo(c, file, i, j, from, to, project)
		}
	}
}
func processSectionOfOrigVideo(c *router.Context, file string,
	i, j, from, to int, project map[string]any) {
	//one := c.One("link", "where guid=$1", guid)
	//c.FreeFormUpdate("update links set duration=$1 where guid=$2", d, guid)
	//c.FreeFormUpdate("update links set photos_ready=true where guid=$1", guid)

	sectionId := fmt.Sprintf("%v_%d_%d", project["id"], i, j)
	oneSection := c.One("link_section", "where section=$1", sectionId)
	//output := fmt.Sprintf("data/%s_%d_%d.mp4", guid, i, j)
	output := fmt.Sprintf(BUCKET+"/%s/orig-video/%d_%d.mp4",
		project["guid"], i, j)

	if len(oneSection) == 0 {

		cmd := exec.Command("ffmpeg",
			"-ss", fmt.Sprintf("%d", from),
			"-i", BUCKET+file,
			"-t", "3",
			"-vf", "scale=1280:720",
			//fmt.Sprintf("%d", to),
			"-c:v", "libx264", "-c:a", "aac", "-y",
			output)
		//cmd := exec.Command("ffmpeg", "-i", "data/"+guid+".mp3",
		//	"-ss", fmt.Sprintf("%d", from), "-to",
		//	fmt.Sprintf("%d", to),
		//	"-y",
		//	output)
		fmt.Println(i, j, from, to)
		cmd.CombinedOutput()
		c.Params = map[string]any{}
		c.Params["project_id"] = project["id"]
		c.Params["section"] = sectionId
		c.Params["minute"] = i
		c.Params["sub"] = j
		c.Params["meta"] = 1
		c.Params["guid"] = util.PseudoUuid()
		c.Insert("link_section")
	}
	oneSection = c.One("link_section", "where section=$1", sectionId)
	//fmt.Println(string(b), err)

	if fileExist(output) == false {
		return
	}

	flac := fmt.Sprintf(BUCKET+"%s/flac/%d_%d.flac",
		project["guid"], i, j)
	if intCheck(oneSection["meta"], 1) {
		cmd := exec.Command("ffmpeg", "-i", output,
			"-b:a", "32k", "-ar", "16000", "-acodec", "flac",
			"-y",
			flac)
		fmt.Println("flac", i, j, from, to)
		cmd.CombinedOutput()
		c.FreeFormUpdate("update link_sections set meta=2 where section=$1", sectionId)
	}
	oneSection = c.One("link_section", "where section=$1", sectionId)

	if intCheck(oneSection["meta"], 2) {
		stt := google.Speech(flac)
		c.FreeFormUpdate("update link_sections set meta=3,stt=$1 where section=$2",
			stt, sectionId)
	}

}
