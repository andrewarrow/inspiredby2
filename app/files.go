package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Files(c *router.Context, second, third string) {
	if second == "upload" && third == "" && c.Method == "POST" {
		handleUpload(c)
		return
	}
	c.NotFound = true
}

func handleUpload(c *router.Context) {
	c.ReadFormValuesIntoParams("file")
	router.SaveMultiFiles(c)
}
