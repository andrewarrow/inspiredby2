package app

import (
	"inspiredby2/video"

	"github.com/andrewarrow/feedback/router"
)

func handlePromptsText(c *router.Context, guid string) {
	c.ReadJsonBodyIntoParams()
	words, _ := c.Params["words-"+guid].(string)
	c.FreeFormUpdate("update link_sections set prompt_text=$1 where guid=$2",
		words, guid)

	go func() {
		//pika.CreateFolder(guid)
		video.AddToPika(c, words, guid)
	}()

	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
