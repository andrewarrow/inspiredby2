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
		count := 0
		for _, item := range items {
			url, _ := item["video_url"].(string)
			guid := "96317f74-01fa-4afb-a681-56a4c607c0c4"
			if url == "" {
				continue
			}
			name := fmt.Sprintf("%03d", count)
			fmt.Println(count, name, url)
			util.Download("data2", "orig_"+name, url)
			convertFrameRate("data2/orig_"+name+".mp4", name)

			copyFile12(fmt.Sprintf("data/%s_%d_%d.mp4", guid, item["minute"], item["sub"]), fmt.Sprintf("data3/%03d.mp4", count))
			count += 1
			time.Sleep(time.Second * 1)
		}
	} else if id == "3" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		other := 1
		orig := 2
		file1 := "data2/000.wav"
		file2 := "data3/000.wav"
		util.RunFF(fmt.Sprintf("-i %s -i %s -filter_complex amix=inputs=2:duration=longest", file1, file2), "data4/000.wav")
		for i, item := range items {
			_ = item
			name := fmt.Sprintf("%03d", i)
			otherName := fmt.Sprintf("%03d", other)
			origName := fmt.Sprintf("%03d", orig)
			//0,v:pika000,a:pike000+orig000
			//1,v:orig001,a:orig001
			//2,v:pika001,a:pike001+orig002
			//3,v:orig003,a:orig003
			//4,v:pika002,a:pika002+orig004
			//5:v:orig005,a:orig005
			if i > 0 && i%2 == 0 {
				file1 := "data2/" + otherName + ".wav"
				file2 := "data3/" + origName + ".wav"
				util.RunFF(fmt.Sprintf("-i %s -i %s -filter_complex amix=inputs=2:duration=longest", file1, file2), "data4/"+name+".wav")
				other++
				orig += 2
			}

		}
	} else if id == "4" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, _ := range items {

			if i%2 == 0 {
				continue
			}
			name := fmt.Sprintf("%03d", i)
			copyFile12("data3/"+name+".wav", "data4/"+name+".wav")
		}
	} else if id == "5" {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		other := 0
		orig := 1
		for i, item := range items {
			_ = item

			//0,v:pika000,a:pike000+orig000
			//1,v:orig001,a:orig001
			//2,v:pika001,a:pike001+orig002
			//3,v:orig003,a:orig003
			//4,v:pika002,a:pika002+orig004
			//5:v:orig005,a:orig005
			name := fmt.Sprintf("%03d", i)
			otherName := fmt.Sprintf("%03d", other)
			origName := fmt.Sprintf("%03d", orig)

			if i%2 == 0 {
				file1 := "data2/" + otherName + ".mp4"
				util.RunFF(fmt.Sprintf("-i %s -c:v copy -an", file1),
					"data5/"+name+".mp4")
				other++
			} else {
				file2 := "data3/" + origName + ".mp4"
				util.RunFF(fmt.Sprintf("-i %s -c:v copy -an", file2),
					"data5/"+name+".mp4")
				orig += 2
			}
		}
	} else if id == "6" {
	} else if id == "7" {
	}
}

func convertFrameRate(filename, name string) {
	//-c:v libx264 -profile:v high -level:v 4.0 -pix_fmt yuv420p -vf scale=1280:720 -r 25 -b:v 509k -c:a copy
	util.RunFF("-i "+filename+" -c:v libx264 -profile:v high -level:v 4.0 -pix_fmt yuv420p -vf scale=1280:720 -r 25 -b:v 509k -c:a copy",
		"data2/"+name+".mp4")
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
