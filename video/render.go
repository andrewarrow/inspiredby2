package video

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func Render(c *router.Context, id int) {
	if id == 1 {
		items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
		for i, item := range items {
			poster := item["video_poster"]
			fmt.Println(i, poster)
		}
	} else if id == 2 {
	}
}
