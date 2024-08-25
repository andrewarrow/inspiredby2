package browser

import (
	"fmt"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

func PollForUpdates() {
	guid := Document.Id("guid").Get("innerHTML")
	time.Sleep(time.Second * 1)

	for {
		json, _ := wasm.DoGet("/core/poll/" + guid)
		fmt.Println(json)

		time.Sleep(time.Millisecond * 3000)
	}
}
