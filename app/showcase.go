package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Showcase(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		//	handleShowcaseIndex(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleShowcaseItem(c, second)
		return
	}
	c.NotFound = true
}
func handleShowcaseItem(c *router.Context, id string) {
	send := map[string]any{}
	c.SendContentInLayout("showcase.html", send, 200)
}
