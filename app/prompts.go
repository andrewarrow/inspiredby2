package app

import (
	"inspiredby2/video"
	"strings"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func Prompts(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		//	handlePromptsIndex(c)
		return
	}
	if second != "" && third == "options" && c.Method == "GET" {
		handlePromptsOptions(c, second)
		return
	}
	if second != "" && third == "text" && c.Method == "POST" {
		handlePromptsText(c, second)
		return
	}
	if second != "" && third == "pika" && c.Method == "DELETE" {
		handlePromptsPikaDelete(c, second)
		return
	}
	if second != "" && third == "hide" && c.Method == "POST" {
		handlePromptsHide(c, second)
		return
	}
	if second != "" && third == "bump" && c.Method == "POST" {
		handlePromptsBump(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handlePromptsItem(c, second)
		return
	}
	c.NotFound = true
}

func handlePromptsItem(c *router.Context, id string) {
	c.Title = "Heart Rate Variability"
	one := c.One("project", "where guid=$1", id)
	items := c.FreeFormSelect("select * from link_sections where project_id=$1 order by minute,sub limit 1000", one["id"])
	send := map[string]any{}

	for _, item := range items {
		stt := item["stt"].(string)
		item["longest"] = strings.Join(video.FindLongestWords(stt), " ")
		item["has_prompt"] = item["prompt_text"] != ""
	}
	send["items"] = items
	c.SendContentInLayout("prompts.html", send, 200)
}

func FixGuids(c *router.Context) {
	items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
	for _, item := range items {
		guid := util.PseudoUuid()
		c.FreeFormUpdate("update link_sections set guid=$1 where id=$2", guid, item["id"])
	}
}
