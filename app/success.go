package app

import "github.com/andrewarrow/feedback/router"

func handleStripeSuccess(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("success.html", send, 200)
}
