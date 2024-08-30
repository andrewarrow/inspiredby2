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
	send := map[string]any{}
	c.SendContentInLayout("prompts.html", send, 200)
}
