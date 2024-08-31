package app

import "github.com/andrewarrow/feedback/router"

func handlePromptsBump(c *router.Context, guid string) {
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
