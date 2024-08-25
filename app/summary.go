package app

import (
	"fmt"
	"inspiredby2/groq"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func groupArrayIntoThree(arr []map[string]any) [][]map[string]any {
	result := make([][]map[string]any, 3)
	itemsPerGroup := len(arr) / 3
	remainingItems := len(arr) % 3

	start := 0
	for i := 0; i < 3; i++ {
		end := start + itemsPerGroup
		if i < remainingItems {
			end++
		}
		result[i] = arr[start:end]
		start = end
	}

	return result
}

func ProcessVideoSummary(c *router.Context, guid string) {
	one := c.One("link", "where guid=$1", guid)

	items := c.All("link_minute", "where link_id=$1 order by minute", "", one["id"])
	summaries := []string{}
	for _, g := range groupArrayIntoThree(items) {
		buffer := []string{}
		for _, item := range g {
			s := item["summary"].(string)
			buffer = append(buffer, s)
		}
		summaries = append(summaries, strings.Join(buffer, " "))
	}
	fmt.Println(len(summaries))
	first := c.One("link_third", "where link_id=$1 and link_key=$2", one["id"], "first")
	second := c.One("link_third", "where link_id=$1 and link_key=$2", one["id"], "second")
	third := c.One("link_third", "where link_id=$1 and link_key=$2", one["id"], "third")

	c.Params = map[string]any{}

	if len(first) == 0 {
		s := groq.Summarize(summaries[0])
		c.Params["link_id"] = one["id"]
		c.Params["link_key"] = "first"
		c.Params["summary"] = s
		c.Insert("link_third")
	}

	if len(second) == 0 {
		s := groq.Summarize(summaries[1])
		c.Params["link_id"] = one["id"]
		c.Params["link_key"] = "second"
		c.Params["summary"] = s
		c.Insert("link_third")
	}

	if len(third) == 0 {
		s := groq.Summarize(summaries[2])
		c.Params["link_id"] = one["id"]
		c.Params["link_key"] = "third"
		c.Params["summary"] = s
		c.Insert("link_third")
	}

}
