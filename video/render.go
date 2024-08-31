package video

import (
	"fmt"
	"inspiredby2/util"
	"os/exec"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func Render(c *router.Context, id string) {
	if id == "1" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			poster, _ := item["video_poster"].(string)
			fmt.Println(i, poster)
			if poster == "" {
				continue
			}
			util.Download("posters", item["guid"].(string), poster)
			time.Sleep(time.Second * 1)
		}
	} else if id == "2" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			url, _ := item["video_url"].(string)
			if url == "" {
				continue
			}
			fmt.Println(i, url)
			util.Download("data2", item["guid"].(string), url)
			time.Sleep(time.Second * 1)
		}
	} else if id == "3" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for _, item := range items {
			guid, _ := item["guid"].(string)
			// make :q
			make12Seconds(guid)

		}
	} else if id == "3" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			guid, _ := item["guid"].(string)

			file1 := fmt.Sprintf("posters/%s.mp4", guid)
			file2 := fmt.Sprintf("data2/%s.mp4", guid)
			file3 := fmt.Sprintf("data3/%03d.mp4", i)
			CombineTwoFilesWithBox(file1, file2, file3)
		}
	}
}

func make12Seconds(guid string) {
	//ffmpeg -loop 1 -i ffebc79d-f388-0686-5fbc-de4e98ed4d16_poster.jpg_720.jpg -c:v libx264 -t 12 -pix_fmt yuv420p -vf "scale=1920:1080" output_video.mp4
	cmd := exec.Command("ffmpeg", "-loop", "1", "-i",
		"posters/"+guid+"_poster.jpg_720.jpg",
		"-c:v", "libx264", "-t", "12",
		"-pix_fmt", "yuv420p", "-vf", "scale=1280:720", "-y", "posters/"+guid+".mp4")
	cmd.CombinedOutput()
	//b, err := cmd.CombinedOutput()
	//fmt.Println(string(b), err)
}
