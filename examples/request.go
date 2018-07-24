package main

import "github.com/sourcegraph/stripe-go/client"

func simpleRequestExample() {
	sc := &client.API{}
	sc.Init("sk_test_BQokikJOvBiI2HlWgH4olfQ2", nil)
	sc.Charges.Get("ch_1CqCgQ2eZvKYlo2CzBnALR20", params)
}
