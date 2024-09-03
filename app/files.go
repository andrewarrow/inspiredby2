package app

import (
	"os"

	"github.com/andrewarrow/feedback/router"
)

func Files(c *router.Context, second, third string) {
	if second == "upload" && third == "" && c.Method == "POST" {
		handleUpload(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleProject(c, second)
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
	guid := c.Params["guid"].(string)

	os.Mkdir(BUCKET+guid, 0755)
	os.Mkdir(BUCKET+guid+"/orig-video", 0755)
	os.Mkdir(BUCKET+guid+"/orig-audio", 0755)
	os.Mkdir(BUCKET+guid+"/pika-video", 0755)
	os.Mkdir(BUCKET+guid+"/pika-audio", 0755)
	os.Mkdir(BUCKET+guid+"/combine", 0755)
	send := map[string]any{}
	c.SendContentAsJson(send, 200)
}
