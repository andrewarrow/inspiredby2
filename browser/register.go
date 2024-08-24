package browser

import (
	"net/url"
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	LogoutEvents()
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/core/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "/core/start")
	}
	if Global.Start == "start.html" {
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "core", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "core", nil, afterRegister)
	} else if Global.Start == "welcome.html" {
		//Global.SubmitEvent("welcome-form", HandleWelcome)
		if HandleWelcomeStep1() == false {
			return
		}
		a := wasm.NewAutoForm("welcome-form")
		a.Path = "/core/add"
		a.Clear = true
		a.Before = func() string {
			Document.Id("go").Set("value", "please wait...")
			return ""
		}
		a.After = func(content string) {
			Global.Location.Set("href", "/core/stripe?email="+url.QueryEscape(content))
		}
		Global.AddAutoForm(a)
	}
}

func HandleWelcomeStep1() bool {
	link := Document.Id("link").Get("value")
	email := Document.Id("email").Get("value")
	if validateEmail(email) != nil {
		Global.Global.Get("alert").Invoke("please enter valid email")
		return false
	}

	if strings.HasPrefix(link, "www") {
		link = "https://" + link
	}

	if strings.HasPrefix(link, "https://www.youtube.com/watch") ||
		strings.HasPrefix(link, "https://youtu.be") {
		return true
	}

	Global.Global.Get("alert").Invoke("please enter valid youtube link")
	return false
}

func LogoutEvents() {
	if Document.Id("logout") == nil {
		return
	}
	Global.Event("logout", Global.Logout("/core"))
}
