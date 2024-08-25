package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

func PollForUpdates() {
	guid := Document.Id("guid").Get("innerHTML")
	guid = strings.TrimSpace(guid)
	fmt.Println(guid)
	time.Sleep(time.Second * 1)

	for {
		json, _ := wasm.DoGet("/core/poll/" + guid)
		fmt.Println(json)

		time.Sleep(time.Millisecond * 3000)
	}
}
