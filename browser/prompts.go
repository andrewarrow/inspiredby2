package browser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/wasm"
)

func SetupPrompts() {
	w := Document.Id("top")
	all := w.SelectAllByClass("prompt-form")

	for _, item := range all {

		guid := item.Id[2:]
		Document.Id("a-" + guid).Event(ClickFetch)
		a := wasm.NewAutoForm(item.Id)
		a.Path = "/prompts/" + guid + "/text"
		a.Clear = true
		a.Before = func() string {
			Document.Id("b-"+guid).Set("value", "please wait...")
			return ""
		}
		a.After = func(content string) {
			Document.Id("b-"+guid).Set("value", "bump")
			handlePromptReply(content)
		}
		Global.AddAutoForm(a)
	}
}

var imageTemplate = `<a href="/" id="delete-%s"><img id="images-%s" src="%s" class="w-64"/></a>`

func ClickFetch(id string) {
	guid := id[2:]
	go func() {
		m := wasm.DoGetMap("/prompts/" + guid + "/options")
		items, _ := m["items"].([]any)
		buffer := []string{}
		pikaIds := []string{}
		for _, item := range items {
			thing, _ := item.(map[string]any)
			buffer = append(buffer, fmt.Sprintf(imageTemplate,
				thing["id_pika"], thing["id_pika"], thing["video_poster"]))
			pikaIds = append(pikaIds, thing["id_pika"].(string))
		}
		join := strings.Join(buffer, "<br/>")
		Document.Id("posters-"+guid).Set("innerHTML", join)
		for _, id := range pikaIds {
			Document.Id("delete-" + id).Event(handleDeletePika)
		}

	}()
}

func handleDeletePika(id string) {
	guid := id[7:]
	go wasm.DoDelete("/prompts/" + guid + "/pika")
}

func handlePromptReply(js string) {
	var m map[string]any
	json.Unmarshal([]byte(js), &m)
	items := m["items"].([]any)
	for _, item := range items {
		thing := item.(map[string]any)

		fmt.Println(thing)
		model := models.NewBase(thing)
		guid := model.GetString("guid")
		duration := model.GetFloatAsInt("duration")
		poster := model.GetString("video_poster")

		if poster != "" {
			w := Document.Id("p-" + guid)
			imgs := w.SelectAllByQuery("getElementsByTagName", "img")
			if len(imgs) > 0 {
				imgs[0].Set("src", poster)
			}
		}

		w := Document.Id("d-" + guid)
		w.Set("innerHTML", fmt.Sprintf("%d", duration))

	}
}
