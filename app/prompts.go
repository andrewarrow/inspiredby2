package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func FixGuids(c *router.Context) {
	items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
	for _, item := range items {
		guid := util.PseudoUuid()
		c.FreeFormUpdate("update link_sections set guid=$1 where id=$2", guid, item["id"])
	}
}

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
	items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
	send := map[string]any{}
	send["items"] = items
	c.SendContentInLayout("prompts.html", send, 200)
}
