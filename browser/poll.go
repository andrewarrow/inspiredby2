package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

var photos bool
var sectionMap map[string]bool = map[string]bool{}

func PollForUpdates() {
	guid := Document.Id("guid").Get("innerHTML")
	guid = strings.TrimSpace(guid)
	fmt.Println(guid)
	time.Sleep(time.Second * 1)

	for {
		m := wasm.DoGetMap("/core/poll/" + guid)

		list := m["all"].([]any)

		if m["photos"] == true {
			for _, item := range list {
				fmt.Println(item)
				thing := item.(map[string]any)
				minute := int(thing["minute"].(float64))
				sub := int(thing["sub"].(float64))
				meta := int(thing["meta"].(float64))
				metaString := ""
				if meta == 1 {
					metaString = "EXTRACTED"
				} else if meta == 2 {
					metaString = "TRANSCRIBED"
				}

				div := Document.Id(fmt.Sprintf("sub-%d-%d", minute, sub))
				div.Set("innerHTML", fmt.Sprintf("&nbsp;&nbsp;section %d - %s",
					sub+1, metaString))
			}
		}

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
			for i := 0; i < minutes+1; i++ {
				inside := `<div>
<div>
minute %d
</div>
<div id="sub-%d-0">
&nbsp;&nbsp;section 1
</div>
<div id="sub-%d-1">
&nbsp;&nbsp;section 2
</div>
<div id="sub-%d-2">
&nbsp;&nbsp;section 3
</div>
<div id="sub-%d-3">
&nbsp;&nbsp;section 4
</div>
<div id="sub-%d-4">
&nbsp;&nbsp;section 5
</div>
<div id="sub-%d-5">
&nbsp;&nbsp;section 6
</div>
</div>`
				div := Document.NewTag("div", fmt.Sprintf(inside, i+1,
					i, i, i, i, i, i))
				div.Set("id", fmt.Sprintf("minute-%d", i))
				canvas.AppendChild(div.JValue)
			}

		}

		time.Sleep(time.Millisecond * 3000)
	}
}
