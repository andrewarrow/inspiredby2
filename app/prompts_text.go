package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
)

func handlePromptsText(c *router.Context, guid string) {
	c.ReadJsonBodyIntoParams()
	fmt.Println(c.Params)
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
