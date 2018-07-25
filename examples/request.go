package main

import (
	stripe "github.com/sourcegraph/stripe-go"
	"github.com/sourcegraph/stripe-go/client"
)

func simpleRequestExample() {
	var params *stripe.ChargeParams

	sc := &client.API{}
	sc.Init("sk_test_BQokikJOvBiI2HlWgH4olfQ2", nil)
	sc.Charges.Get("ch_1CqCgQ2eZvKYlo2CzBnALR20", params)
}
