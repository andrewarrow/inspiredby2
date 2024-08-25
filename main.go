package main

import (
	"embed"
	"inspiredby2/app"
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
	} else if arg == "video" {
		fr := router.NewRouter("DATABASE_URL", embeddedFile)
		c := fr.ToContext()
		app.ProcessVideo(c, os.Args[2])
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		r.Paths["core"] = app.Core
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
