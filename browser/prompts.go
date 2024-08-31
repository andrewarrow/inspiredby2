package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

func SetupPrompts() {
	w := Document.Id("top")
	all := w.SelectAllByClass("prompt-form")

	for _, item := range all {

		guid := item.Id[2:]
		a := wasm.NewAutoForm(item.Id)
		a.Path = "/promprts/" + guid + "/" + bump
		a.Clear = true
		a.Before = func() string {
			Document.Id("b-"+guid).Set("value", "please wait...")
			return ""
		}
		a.After = func(content string) {
			Document.Id("b-"+guid).Set("value", "bump")
		}
		Global.AddAutoForm(a)
	}
}
