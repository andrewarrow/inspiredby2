package app

import (
	"os"

	"github.com/andrewarrow/feedback/router"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func handleProfileStripePost(c *router.Context) {
	email := c.Request.URL.Query().Get("email")
	domain := os.Getenv("LINK_DOMAIN")
	returnPath := "/"
	params := &stripe.CheckoutSessionParams{
		AllowPromotionCodes: stripe.Bool(true),
		CustomerEmail:       stripe.String(email),
		Mode:                stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1PrSdWHFkhRYMfGpU3IdtnJS"),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(domain + "/core/success"),
		CancelURL:  stripe.String(domain + returnPath),
	}

	s, err := session.New(params)

	if err != nil {
		router.SetFlash(c, err.Error())
		router.Redirect(c, "/")
		return
	}

	c.FreeFormUpdate("update users set id_stripe_session=$1 where email=$2", s.ID, email)
	router.SetCookie(c, "id_stripe_session", s.ID)

	router.Redirect(c, s.URL)
}
