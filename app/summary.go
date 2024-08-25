package app

import (
	"fmt"
	"inspiredby2/groq"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func ProcessVideoSummary(c *router.Context, guid string) {
	one := c.One("link", "where guid=$1", guid)

	minuteKey := fmt.Sprintf("%d_%d", one["id"], i)
	oneMinute := c.One("link_minute", "where minute_key=$1", minuteKey)

	if len(oneMinute) == 0 {

		all := c.All("link_section", "where link_id=$1 and minute=$2 order by sub", "",
			one["id"], i)
		buffer := []string{}
		for _, item := range all {
			buffer = append(buffer, item["stt"].(string))
		}
		s := groq.Summarize(strings.Join(buffer, " "))
		c.Params = map[string]any{}
		c.Params["link_id"] = one["id"]
		c.Params["minute"] = i
		c.Params["minute_key"] = minuteKey
		c.Params["summary"] = s
		c.Insert("link_minute")
	}
}
