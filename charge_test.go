package coinbase

import (
	"fmt"
	"testing"
)

func TestCharges(t *testing.T) {
	c := NewTestClient()
	charges, err := c.ListCharges()
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	fmt.Println(len(charges))
	return

}

func TestRetrieveCharge(t *testing.T) {
	c := NewTestClient()
	_, err := c.RetrieveCharge("76936406-dac7-45f9-8efd-b7a5f8efa7ee")
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

func TestCreateCharge(t *testing.T) {
	c := NewTestClient()

	chargeRequest := ChargeRequest{
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		PricingType: FixedPrice,
		LocalPrice: LocalPrice{
			Amount:   "100.00",
			Currency: "USD",
		},
	}

	charge, err := c.CreateCharge(chargeRequest)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(charge)
}

func TestCancelCharge(t *testing.T) {
	c := NewTestClient()

	chargeRequest := ChargeRequest{
		Name:        "The Sovereign Individual",
		Description: "Mastering the Transition to the Information Age",
		PricingType: FixedPrice,
		LocalPrice: LocalPrice{
			Amount:   "100.00",
			Currency: "USD",
		},
	}

	charge, err := c.CreateCharge(chargeRequest)
	if err != nil {
		t.Fatal(err)
	}

	returnedCharge, err := c.CancelCharge(charge.ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(returnedCharge)
}
