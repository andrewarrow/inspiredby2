package app

import "github.com/andrewarrow/feedback/router"

func handleDemo(c *router.Context) {
	link := map[string]any{}
	link["link"] = "https://youtu.be/wh_M25S2xUw"
	send := map[string]any{}
	send["item"] = link
	c.SendContentInLayout("demo.html", send, 200)
}
