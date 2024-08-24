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
		Global.SubmitEvent("welcome-form", HandleWelcome)
	}
}

func HandleWelcome() {
	link := Document.Id("link").Get("value")
	email := Document.Id("email").Get("value")
	if validateEmail(email) != nil {
		Global.Global.Get("alert").Invoke("please enter valid email")
		return
	}

	if strings.HasPrefix(link, "https://www.youtube.com/watch") ||
		strings.HasPrefix(link, "https://youtu.be") {
		Global.Location.Set("href", "https://buy.stripe.com/test_cN23e40qW2fA024cMM?prefilled_email="+url.QueryEscape(email))
		return
	}

	Global.Global.Get("alert").Invoke("please enter valid youtube link")
}

func LogoutEvents() {
	if Document.Id("logout") == nil {
		return
	}
	Global.Event("logout", Global.Logout("/core"))
}
