package video

import (
	"fmt"
	"inspiredby2/util"
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
	}
}
