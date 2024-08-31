package video

import (
	"fmt"
	"inspiredby2/util"

	"github.com/andrewarrow/feedback/router"
)

func Render(c *router.Context, id int) {
	if id == 1 {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			poster, _ := item["video_poster"].(string)
			fmt.Println(i, poster)
			if poster == "" {
				continue
			}
			util.Download("posters", item["guid"].(string), poster)
		}
	} else if id == 2 {
	}
}
