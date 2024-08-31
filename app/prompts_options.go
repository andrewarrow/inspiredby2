package app

import "github.com/andrewarrow/feedback/router"

func handlePromptsOptions(c *router.Context, guid string) {
	send := map[string]any{}
	one := c.One("link_section", "where guid=$1", guid)
	all := c.All("pika", "where link_section_id=$1", "", one["id"])
	send["items"] = all
	c.SendContentAsJson(send, 200)
}
