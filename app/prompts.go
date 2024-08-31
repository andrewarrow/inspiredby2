package app

import (
	"sort"
	"strings"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func Prompts(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		//	handlePromptsIndex(c)
		return
	}
	if second != "" && third == "text" && c.Method == "POST" {
		handlePromptsText(c, second)
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
	items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
	send := map[string]any{}

	for _, item := range items {
		stt := item["stt"].(string)
		item["longest"] = strings.Join(findLongestWords(stt), " ")
	}
	send["items"] = items
	c.SendContentInLayout("prompts.html", send, 200)
}

func findLongestWords(input string) []string {
	words := strings.Split(input, " ")

	for i, word := range words {
		words[i] = strings.Trim(word, ".,!?\"'`")
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	if len(words) >= 3 {
		return words[:3]
	}
	return words
}

func FixGuids(c *router.Context) {
	items := c.FreeFormSelect("select * from link_sections order by minute,sub limit 1000")
	for _, item := range items {
		guid := util.PseudoUuid()
		c.FreeFormUpdate("update link_sections set guid=$1 where id=$2", guid, item["id"])
	}
}
