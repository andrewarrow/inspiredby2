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
		items := c.FreeFormSelect("select * from link_sections where meta!=90 and minute < 2 order by minute,sub limit 1000")
		count := 0
		for _, item := range items {
			url, _ := item["video_url"].(string)
			if url == "" {
				continue
			}
			fmt.Println(count, fmt.Sprintf("%03d", count), url)
			util.Download("data3", fmt.Sprintf("%03d", count), url)
			count += 2
			time.Sleep(time.Second * 1)
		}
	} else if id == "3" {
		count := 0
		items := c.FreeFormSelect("select * from link_sections where meta!=90 and minute < 2 order by minute,sub limit 1000")
		for _, item := range items {
			url, _ := item["video_url"].(string)
			if url == "" {
				continue
			}
			name := fmt.Sprintf("%03d", count)
			//  ffmpeg -i "$f" -acodec pcm_s16le -ar 44100 -ac 2 "wav_files/${f%.mp3}.wav"
			util.RunFF("-i data3/"+name+".mp4 "+
				"-t 3 -acodec pcm_s16le -ar 44100 -ac 2",
				"data4/"+name+".wav")
			count += 2

		}
	} else if id == "4" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			guid, _ := item["guid"].(string)

			file1 := fmt.Sprintf("posters/%s.mp4", guid)
			file2 := fmt.Sprintf("data2/%s.mp4", guid)
			file3 := fmt.Sprintf("data3/%03d.mp4", i)
			CombineTwoFilesWithBox(file1, file2, file3)
			changeToMov(fmt.Sprintf("%03d", i), file3)
		}
	} else if id == "5" {
		Combine("data3")
	} else if id == "6" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			fmt.Println("wefwef", item["id"])
			one := c.One("pika_render", "where link_section_id=$1 and duration=15", item["id"])
			url, _ := one["video_url"].(string)
			if url == "" {
				continue
			}
			fmt.Println(i, url)
			util.Download("data4", fmt.Sprintf("%03d", i), url)
			time.Sleep(time.Second * 1)
		}
	} else if id == "7" {
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
func changeToMov(guid, file string) {
	cmd := exec.Command("ffmpeg", "-i", file,
		"-y", "data3/"+guid+".mov")
	cmd.CombinedOutput()
	//b, err := cmd.CombinedOutput()
	//fmt.Println(string(b), err)
}
