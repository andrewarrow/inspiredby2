package app

import (
	"os"

	"github.com/andrewarrow/feedback/router"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func handleProfileStripePost(c *router.Context) {
	email, ok := c.User["email"].(string)
	if !ok {
		router.SetFlash(c, "set email first")
		router.Redirect(c, "/")
		return
	}
	domain := os.Getenv("LINK_DOMAIN")
	returnPath := "/"
	params := &stripe.CheckoutSessionParams{
		AllowPromotionCodes: stripe.Bool(true),
		CustomerEmail:       stripe.String(email),
		Mode:                stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("prod_QiuSB6v8X69FXs"),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(domain + "/customer/success"),
		CancelURL:  stripe.String(domain + returnPath),
	}

	s, err := session.New(params)

	if err != nil {
		router.SetFlash(c, err.Error())
		router.Redirect(c, "/")
		return
	}

	//c.FreeFormUpdate("update users set id_stripe_session=$1 where id=$2", s.ID, c.User["id"])

	router.Redirect(c, s.URL)
}
