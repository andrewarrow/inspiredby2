package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func handlePoll(c *router.Context, guid string) {
	m := map[string]any{}
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
	one := c.One("user", "where id_stripe_session=$1", sid)
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
