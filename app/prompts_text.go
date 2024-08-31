package app

import (
	"inspiredby2/pika"

	"github.com/andrewarrow/feedback/router"
)

var flavors = []string{"3D render", "Moody, filmic style, 35mm", ""}

func handlePromptsText(c *router.Context, guid string) {
	c.ReadJsonBodyIntoParams()
	words, _ := c.Params["words-"+guid].(string)
	c.FreeFormUpdate("update link_sections set prompt_text=$1 where guid=$2",
		words, guid)

	go func() {
		one := c.One("link_section", "where guid=$1", guid)
		for _, flavor := range flavors {
			pt := flavor + " " + words
			id := pika.Generate("", pt)
			c.Params = map[string]any{}
			c.Params["id_pika"] = id
			c.Params["prompt_text"] = pt
			c.Params["link_section_id"] = one["id"]
			c.Insert("pika")
		}
	}()

	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
