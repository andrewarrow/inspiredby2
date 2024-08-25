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
		m := wasm.DoGetMap("/core/poll/" + guid)
		if m["photos"] == true {
			p1 := Document.Id("photo1")
			p2 := Document.Id("photo2")
			p1.Set("src", "/bucket/"+guid+"_1.jpg")
			p2.Set("src", "/bucket/"+guid+"_2.jpg")
			duration := Document.Id("duration")
			d, _ := m["duration"].(string)
			duration.Set("innerHTML", FormatSeconds(d))
		}

		time.Sleep(time.Millisecond * 3000)
	}
}
