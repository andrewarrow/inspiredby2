package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Prompts(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		//	handlePromptsIndex(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handlePromptsItem(c, second)
		return
	}
	c.NotFound = true
}
func handlePromptsItem(c *router.Context, id string) {
	c.Title = "Heart Rate Variability"
	items := c.FreeFormSelect("select * from link_sections order by section limit 1000")
	send := map[string]any{}
	send["items"] = items
	c.SendContentInLayout("prompts.html", send, 200)
}
