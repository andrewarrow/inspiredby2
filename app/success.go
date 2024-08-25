package app

import (
	"fmt"
	"time"

	"github.com/andrewarrow/feedback/router"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func handlePoll(c *router.Context, guid string) {
	m := map[string]any{}
	one := c.One("link", "where guid=$1", guid)
	if one["photos_ready"] == "1" {
		m["photos"] = true
		m["duration"] = one["duration"]
	}
	all := c.FreeFormSelect("select minute,sub,meta from link_sections where link_id=$1", one["id"])
	m["all"] = all

	all = c.FreeFormSelect("select minute,summary from link_minutes where link_id=$1",
		one["id"])
	m["summary"] = all
	c.SendContentAsJson(m, 200)
}
func handleAdd(c *router.Context) {
	c.ReadJsonBodyIntoParams()
	c.ValidateCreate("user")
	c.Insert("user")
	m := map[string]any{"guid": c.Params["guid"]}
	c.SendContentAsJson(m, 200)
}

func handleStripeSuccess(c *router.Context) {
	sid := router.GetCookie(c, "id_stripe_session")
	one := c.One("user", "where id_stripe_session=$1", sid)
	if one["verified_at"] == nil {
		params := &stripe.CheckoutSessionParams{}
		params.AddExpand("payment_intent")
		params.AddExpand("line_items")
		params.AddExpand("line_items.data")
		params.AddExpand("line_items.data.price")
		checkoutSession, err := session.Get(sid, params)
		if err != nil {
			fmt.Println(err)
			router.Redirect(c, "/")
			return
		}
		if checkoutSession.PaymentStatus != "paid" {
			router.Redirect(c, "/")
			return
		}
		c.FreeFormUpdate("update users set verified_at=$1 where id_stripe_session=$2",
			time.Now().Unix(), sid)
	}
	link := c.One("link", "where link=$1", one["link"])
	if len(link) == 0 {
		c.Params = map[string]any{}
		c.Params["link"] = one["link"]
		c.Params["user_id"] = one["id"]
		c.ValidateCreate("link")
		c.Insert("link")
	}
	link = c.One("link", "where link=$1", one["link"])
	c.SendContentInLayout("success.html", link, 200)
}
