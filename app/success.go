package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func handleAdd(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	fmt.Println(c.Params)

	c.ValidateCreate("user")
	c.Insert("user")

	m := map[string]any{"guid": c.Params["guid"]}

	c.SendContentAsJson(m, 200)
}
func handleStripeSuccess(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("success.html", send, 200)
}
