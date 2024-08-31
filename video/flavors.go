package video

import (
	"inspiredby2/pika"
	"sort"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

var Flavors = []string{"3D render", "Moody, filmic style, 35mm", ""}

func AddToPika(c *router.Context, words, guid string) {
	one := c.One("link_section", "where guid=$1", guid)
	for _, flavor := range Flavors {
		pt := flavor + " " + words
		id := pika.Generate("", pt)
		if id == "" {
			continue
		}
		//pika.MoveVideoToFolder(id, "A31")
		c.Params = map[string]any{}
		c.Params["id_pika"] = id
		c.Params["prompt_text"] = pt
		c.Params["link_section_id"] = one["id"]
		c.Insert("pika")
	}
}
func FindLongestWords(input string) []string {
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
