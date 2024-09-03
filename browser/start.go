package browser

import (
	"fmt"
	"syscall/js"
)

func SetupStart() {
	Global.SubmitEvent("upload-form", handleUpload)
}

func handleUpload() {
	w := Document.Id("file")
	list := w.JValue.Get("files")
	if list.Length() == 0 {
		Global.Global.Call("alert", "Please select a file.")
		return
	}
	file := list.Index(0)
	name := Document.Id("name").Get("value")

	formData := Global.Global.Get("FormData").New()
	formData.Call("append", "file", file)
	formData.Call("append", "name", name)

	xhr := Global.Global.Get("XMLHttpRequest").New()
	xhr.Get("upload").Call("addEventListener", "progress", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		fmt.Println(event)
		if event.Get("lengthComputable").Bool() {
			loaded := event.Get("loaded").Float()
			total := event.Get("total").Float()
			percentComplete := (loaded / total) * 100
			fmt.Println(percentComplete, loaded, total)
		}
		return nil
	}))

	go func() {
		xhr.Call("open", "POST", "/files/upload")
		xhr.Call("send", formData)
	}()

}
