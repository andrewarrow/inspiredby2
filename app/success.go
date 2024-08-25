package app

import (
	"fmt"

	"github.com/andrewarrow/feedback/router"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

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
	c.SendContentInLayout("success.html", one, 200)
}
