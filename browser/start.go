package browser

import (
	"fmt"
	"syscall/js"
)

func SetupStart() {
	Global.SubmitEvent("upload-form", handleUpload)
}

func handleUpload() {
	file := Document.Id("file")
	list := file.JValue.Get("files")
	if list.Length() == 0 {
		Global.Global.Call("alert", "Please select a file.")
		return
	}
	file := fileList.Index(0)

	formData := Global.Global.Get("FormData").New()
	formData.Call("append", "file", file)

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

	xhr.Call("open", "POST", "/files/upload")
	xhr.Call("send", formData)

}
