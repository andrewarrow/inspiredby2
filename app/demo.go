package app

import "github.com/andrewarrow/feedback/router"

func handleDemo(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("demo.html", send, 200)
}
