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
	c.ReadFormValuesIntoParams("file", "name")
	router.SaveMultiFiles(c)
	c.Params["file"] = c.Params["photo"]
	c.Params["user_id"] = c.User["id"]
	c.ValidateCreate("project")
	c.Insert("project")
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
