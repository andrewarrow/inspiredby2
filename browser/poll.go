package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

var photos bool

func PollForUpdates() {
	guid := Document.Id("guid").Get("innerHTML")
	guid = strings.TrimSpace(guid)
	fmt.Println(guid)
	time.Sleep(time.Second * 1)

	for {
		m := wasm.DoGetMap("/core/poll/" + guid)
		if m["photos"] == true && photos == false {
			photos = true
			p1 := Document.Id("photo1")
			p2 := Document.Id("photo2")
			p1.Set("src", "/bucket/"+guid+"_1.jpg")
			p2.Set("src", "/bucket/"+guid+"_2.jpg")
			duration := Document.Id("duration")
			d, _ := m["duration"].(string)
			formatted, minutes := FormatSeconds(d)
			duration.Set("innerHTML", formatted)
			canvas := Document.Id("canvas")
			for i := 0; i < minutes; i++ {
				div := Document.NewTag("div", fmt.Sprintf("minute %d", i+1))
				div.Set("id", fmt.Sprintf("minute-%d", i+1))
				canvas.AppendChild(div.JValue)
			}

		}

		time.Sleep(time.Millisecond * 3000)
	}
}
