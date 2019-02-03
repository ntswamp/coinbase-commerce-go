package coinbase

import (
	"fmt"
	"testing"
)

func TestCheckouts(t *testing.T) {
	c := NewTestClient()
	checkouts, err := c.ListCheckouts()
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	fmt.Println(len(checkouts))
	return

}

func TestRetrieveCheckout(t *testing.T) {
	c := NewTestClient()
	_, err := c.RetrieveCheckout("76936406-dac7-45f9-8efd-b7a5f8efa7ee")
	if err == nil {
		t.Fatal("we should not end up here")
	}

	if err != nil {
		switch err.(type) {
		case ResponseError:
			responseError, _ := err.(ResponseError)
			if !responseError.IsNotFound() || responseError.HttpStatusCode != 404 {
				t.Fatalf("expected the error to be not found, instead received: %+v\n", responseError)
			}
		default:
			t.Fatal("unexpected error: ", err)
		}
	}

}

func TestCreateCheckout(t *testing.T) {
	c := NewTestClient()

	checkoutData := Checkout{
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		PricingType: FixedPrice,
		LocalPrice: LocalPrice{
			Amount:   "100.00",
			Currency: "USD",
		},
		RequestedInfo: []string{"name", "email"},
	}

	checkout, err := c.CreateCheckout(checkoutData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(checkout)
}

func TestUpdateCheckout(t *testing.T) {
	c := NewTestClient()

	checkoutData := Checkout{
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		PricingType: FixedPrice,
		LocalPrice: LocalPrice{
			Amount:   "100.00",
			Currency: "USD",
		},
		RequestedInfo: []string{"name", "email"},
	}

	checkout, err := c.CreateCheckout(checkoutData)
	if err != nil {
		t.Fatal(err)
	}

	checkout.LocalPrice = LocalPrice{
		Amount:   "200",
		Currency: "USD",
	}

	updatedCheckout, err := c.UpdateCheckout(checkout)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(updatedCheckout)
}

func TestDeleteCheckout(t *testing.T) {
	c := NewTestClient()

	checkoutData := Checkout{
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		PricingType: FixedPrice,
		LocalPrice: LocalPrice{
			Amount:   "100.00",
			Currency: "USD",
		},
		RequestedInfo: []string{"name", "email"},
	}

	checkout, err := c.CreateCheckout(checkoutData)
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteCheckout(checkout.ID)

	if err != nil {
		t.Fatal(err)
	}

}
