package browser

import (
	"encoding/json"
	"fmt"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/wasm"
)

func SetupPrompts() {
	w := Document.Id("top")
	all := w.SelectAllByClass("prompt-form")

	for _, item := range all {

		guid := item.Id[2:]
		a := wasm.NewAutoForm(item.Id)
		a.Path = "/prompts/" + guid + "/bump"
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

func handlePromptReply(js string) {
	var m map[string]any
	err := json.Unmarshal([]byte(js), &m)
	items := m["items"].([]any)
	for _, item := range items {
		thing := item.(map[string]any)

		fmt.Println(thing)
		model := models.NewBase(thing)
		guid := model.GetString("guid")
		fmt.Println(guid)
		duration := model.GetFloatAsInt("duration")
		fmt.Println(duration)
	}
}
