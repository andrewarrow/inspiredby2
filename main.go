package main

import (
	"embed"
	"fmt"
	"inspiredby2/app"
	"inspiredby2/google"
	"inspiredby2/groq"
	"inspiredby2/pika"
	"inspiredby2/video"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
	"github.com/stripe/stripe-go/v74"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]
	router.DB_FLAVOR = "sqlite"
	stripe.Key = os.Getenv("STRIPE_SK")

	if arg == "import" {
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "google" {
		google.Speech("")
	} else if arg == "groq" {
		textToSummarize := ` person helping them be aware what it feels like for them at the lower HRV and
what it feels like at higher HIV and then understanding the factors that they can control um to continue to amplify their HRV
talk to me about the relationship between HRV autonomic flexibility and the brain and brain function  so we used to think of
what it feels like at higher HIV and then understanding the factors that they can control um to continue to amplify their HRV
talk to me about the relationship between HRV autonomic flexibility and the brain and brain function  so we used to think of
the heart rate variability just as a metric of the autonomic nervous systems resilience to flexibly respond which means if I need to amp up and run
across the street I can this isn't this isn't associated with relaxation or just being calm it's acclimating or adapting flexibly to the needs of the most
and resilience in the last 10 to 15 years uh myself and several other researchers have focused on heart rate variabilities impact.`
		s := groq.Summarize(textToSummarize)
		fmt.Println(s)
	} else if arg == "fix" {
		app.ProcessVideoFix()
	} else if arg == "demo" {
		video.Demo()
	} else if arg == "splice" {
		video.Splice("data2")
	} else if arg == "remove_bottom" {
		video.RemoveBottom(os.Args[2])
	} else if arg == "Resize1280x720" {
		video.Resize1280x720(os.Args[2])
	} else if arg == "Combine" {
		video.Combine("data5/foo3")
	} else if arg == "pika" {
		pika.FindPrompts()
	} else if arg == "PikaDelete" {
		lastId := ""
		for {
			items, ok := pika.List(lastId)
			if ok == false {
				continue
			}
			fmt.Println(items)
			for _, item := range items {
				pika.Delete(item)
				time.Sleep(time.Second * 1)
			}
			if len(items) == 0 {
				return
			}
			lastId = items[len(items)-1]
		}
	} else if arg == "PikaGenerate" {
		pika.Generate("properly gotten treatment worked")
	} else if arg == "thumb" {
		fr := router.NewRouter("DATABASE_URL", embeddedFile)
		c := fr.ToContext()
		app.ProcessThumbs(c, os.Args[2])
	} else if arg == "video" {
		fr := router.NewRouter("DATABASE_URL", embeddedFile)
		c := fr.ToContext()
		app.ProcessVideo(c, os.Args[2])
	} else if arg == "summary" {
		fr := router.NewRouter("DATABASE_URL", embeddedFile)
		c := fr.ToContext()
		app.ProcessVideoSummary(c, os.Args[2])
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		r.Paths["core"] = app.Core
		r.Paths["showcase"] = app.Showcase
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}
}
