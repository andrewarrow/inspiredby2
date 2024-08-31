package app

import (
	"inspiredby2/pika"

	"github.com/andrewarrow/feedback/router"
)

func handlePromptsBump(c *router.Context, guid string) {
	p1 := map[string]any{}
	p1["guid"] = "da777a8b-8443-391e-9f76-fb921403c022"
	p1["duration"] = 6
	p2 := map[string]any{}
	p2["guid"] = "64cab908-d625-54a4-8d0e-dc85721610f0"
	p2["video_poster"] = "https://cdn.pika.art/v1/3be270ab-697d-4716-8900-e7aa778e3f97/poster.jpg"
	p2["duration"] = 0
	send := map[string]any{}
	items := []any{p1, p2}
	send["items"] = items
	c.SendContentAsJson(send, 200)
}

func handlePromptsPikaDelete(c *router.Context, guid string) {
	send := map[string]any{}
	go pika.Delete(guid)
	c.FreeFormUpdate("delete from pikas where id_pika=$1", guid)
	c.SendContentAsJson(send, 200)
}
