package browser

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

var photos bool
var sectionMap map[string]bool = map[string]bool{}

var jump = false

func jumpFunc() {
	jump = true
}

func PollForDemoUpdates() {
	Global.Event("jump", jumpFunc)
	time.Sleep(time.Second * 1)

	progress := 0
	p := Document.Id("progress")
	jumpLink := Document.Id("jump")

	for {
		p.Set("value", fmt.Sprintf("%d", progress))
		progress += 10
		time.Sleep(time.Millisecond * 1)

		if progress > 100 {
			progress = 0
			p.Set("value", "0")
			break
		}
	}
	msg := Document.Id("msg")
	msg.Set("innerHTML", "Breaking 55 min video into parts.")

	jumpLink.RemoveClass("hidden")
	Document.Id("thumbs").RemoveClass("hidden")
	minutes := 55
	drawEachMinute(minutes)
	p.Set("value", "1")
	for i := 0; i < minutes+1; i++ {
		if jump {
			break
		}
		for j := 0; j < 6; j++ {
			if jump {
				break
			}
			div := Document.Id(fmt.Sprintf("sub-%d-%d", i, j))
			m := wasm.DoGetMap(fmt.Sprintf("/core/demo-poll?key=%d_%d", i, j))
			stt, _ := m["stt"].(string)
			div.Set("innerHTML", fmt.Sprintf("&nbsp;&nbsp;section %d - %s",
				j+1, stt))
			time.Sleep(time.Second * 1)
		}
		progress += 1
		p.Set("value", fmt.Sprintf("%d", progress))
	}

	canvas := Document.Id("canvas")
	canvas.Set("innerHTML", "")
	m := wasm.DoGetMap("/core/demo-poll?key=summaries")
	drawEachMinuteNoSubs(minutes)
	for k, v := range m {
		div := Document.Id(fmt.Sprintf("sub-%s-summary", k))
		div.Set("innerHTML", v)
	}
	msg.Set("innerHTML", "foo.")
}

func drawEachMinuteNoSubs(minutes int) {
	canvas := Document.Id("canvas")
	for i := 0; i < minutes+1; i++ {
		inside := `<div>
<div>
minute %d summary
</div>
<div id="sub-%d-summary" class="bg-grey-100 rounded-lg p-3">
</div>
</div>`
		div := Document.NewTag("div", fmt.Sprintf(inside, i+1,
			i))
		div.Set("id", fmt.Sprintf("minute-%d", i))
		canvas.AppendChild(div.JValue)
	}
	canvas.RemoveClass("hidden")
}
func drawEachMinute(minutes int) {
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
<div id="sub-%d-summary" class="bg-purple-900 rounded-lg p-3 hidden">
</div>
</div>`
		div := Document.NewTag("div", fmt.Sprintf(inside, i+1,
			i, i, i, i, i, i, i))
		div.Set("id", fmt.Sprintf("minute-%d", i))
		canvas.AppendChild(div.JValue)
	}
	canvas.RemoveClass("hidden")
}

func PollForUpdates() {
	guid := Document.Id("guid").Get("innerHTML")
	guid = strings.TrimSpace(guid)
	fmt.Println(guid)
	time.Sleep(time.Second * 1)

	for {
		m := wasm.DoGetMap("/core/poll/" + guid)

		list := m["all"].([]any)
		summary := m["summary"].([]any)

		if photos {
			for _, item := range summary {
				thing := item.(map[string]any)
				summary := thing["summary"].(string)
				minute := int(thing["minute"].(float64))

				div := Document.Id(fmt.Sprintf("sub-%d-summary", minute))
				div.Set("innerHTML", summary)
				div.RemoveClass("hidden")
			}
		}

		if photos {
			for _, item := range list {
				thing := item.(map[string]any)
				minute := int(thing["minute"].(float64))
				sub := int(thing["sub"].(float64))
				meta := int(thing["meta"].(float64))
				metaString := ""
				if meta == 1 {
					metaString = "VIDEO_EXTRACTED"
				} else if meta == 2 {
					metaString = "AUDIO_EXTRACTED"
				} else if meta == 3 {
					metaString = "TRANSCRIBED"
				}

				div := Document.Id(fmt.Sprintf("sub-%d-%d", minute, sub))
				div.Set("innerHTML", fmt.Sprintf("&nbsp;&nbsp;section %d - %s",
					sub+1, metaString))
			}
		}

		if m["photos"] == true && photos == false {
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
<div id="sub-%d-summary" class="bg-purple-900 rounded-lg p-3 hidden">
</div>
</div>`
				div := Document.NewTag("div", fmt.Sprintf(inside, i+1,
					i, i, i, i, i, i, i))
				div.Set("id", fmt.Sprintf("minute-%d", i))
				canvas.AppendChild(div.JValue)
			}
			photos = true

		}

		time.Sleep(time.Millisecond * 3000)
	}
}
