package app

import "github.com/andrewarrow/feedback/router"

func handleProject(c *router.Context, guid string) {
	send := map[string]any{}
	item := c.One("project", "where guid=$1", guid)
	send["item"] = item
	c.SendContentInLayout("project.html", send, 200)
}
