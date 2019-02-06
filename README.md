# Coinbase Commerce Go

The unofficial Go library for the [Coinbase Commerce API](https://commerce.coinbase.com/docs/).

# Table of contents

<!--ts-->
   * [Documentation](#documentation)
   * [Installation](#installation)
   * [Usage](#usage)
      * [Checkouts](#checkouts)
      * [Charges](#charges)
<!--te-->

## Documentation
For more details visit [Coinbase API docs](https://commerce.coinbase.com/docs/api/).

To start using this library register an account on [Coinbase Commerce](https://commerce.coinbase.com/signup).
You will find your ``API_KEY`` from User Settings.

Next initialize a ``HttpCoinbaseClient`` for interacting with the API. The only required parameter to initialize a client is `API_KEY`

**Default timeout for the http client is 3 seconds***, you can change that value like so:

``` go
import "github.com/opaolini/coinbase-commerce-go"

func main() {
    c := coinbase.NewHttpClient("API_KEY").WithClientTimeout(3000) // Optional timeout specification in milliseconds
}

```

## Error handling
Client supports the handling of common API errors and warnings.


| Error                    | Status Code |
|--------------------------|-------------|
| APIError                 |      *      |
| InvalidRequestError      |     400     |
| ParamRequiredError       |     400     |
| ValidationError          |     400     |
| AuthenticationError      |     401     |
| ResourceNotFoundError    |     404     |
| RateLimitExceededError   |     429     |
| InternalServerError      |     500     |
| ServiceUnavailableError  |     503     |

Example of handling errors from the API:

``` go
checkoutID := "random-id"
result, err := c.RetrieveCheckout(checkoutID)
if err != nil {
    switch err.(type) {
    case ResponseError:
        responseError, _ := err.(ResponseError)
        if responseError.IsValidationError() {
            log.Printf("Provided checkout ID: %s is invalid: %s \n", checkoutID, responseError.ReturnedError.Message)
        }
    }
}
```

## Installation

Install with ``go``:
``` sh
go get github.com/opaolini/coinbase-commerce-go
```
## Usage
``` go
import "github.com/opaolini/coinbase-commerce-go"

func main() {
    c := coinbase.NewHttpClient("API_KEY")
}

```
## Checkouts 
[Checkouts API docs](https://commerce.coinbase.com/docs/api/#checkouts)

### Retrieve
``` go
checkout, err := c.RetrieveCheckout(<checkout_id>)
if err != nil {
   panic(err)
}
```

### Create
``` go
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
    panic(err)
}

fmt.Println(checkout)
```
### Update
``` go
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
    panic(err)
}

checkout.LocalPrice = LocalPrice{
    Amount:   "200",
    Currency: "USD",
}

updatedCheckout, err := c.UpdateCheckout(checkout)
if err != nil {
    panic(err)
}

fmt.Println(updatedCheckout)
```
### Delete
``` go
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
    panic(err)
}

err = c.DeleteCheckout(checkout.ID)

if err != nil {
    panic(err)
}
```
### List
``` go
checkouts, err := c.ListCheckouts(nil) // grabs all
if err != nil {
    panic("unexpected error: ", err)
}

fmt.Println(len(checkouts))
```

#### With Pagination
``` go
p := &Pagination{
		Limit: 50,
	}
checkouts, err := c.ListCheckouts(p) 
if err != nil {
    panic("unexpected error: ", err)
}

fmt.Println(len(checkouts))
```

## Charges
[Charges API docs](https://commerce.coinbase.com/docs/api/#charges)

### Retrieve
``` go
charge, err := c.RetrieveCharge("76936406-dac7-45f9-8efd-b7a5f8efa7ee")
if err == nil {
    panic("we should not end up here")
}
fmt.Println(charge)
```
### Create
``` go
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
    panic(err)
}
fmt.Println(charge)
```
### List
``` go
charges, err := c.ListCharges(nil)
if err != nil {
    panic(err)
}

fmt.Println(len(charges))
```

#### With pagination
``` go
p := &Pagination{
		Limit: 50,
	}
charges, err := c.ListCharges(p)
if err != nil {
    panic(err)
}

fmt.Println(len(charges))

```

## Webhooks

### Verify Signature from http.Request
``` go
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	sharedKey := "your-shared-key"
	ok, err := VerifyWebhookSignatureFromRequest(sharedKey, r)
	if err != nil {
		// handle error
		panic(err)
	}

	if !ok {
		// handle invalid signature
	}
}
```


### Verify Signature

``` go
    sharedKey := "your-shared-key"
    signature := ""
    body := []byte("{}")
    ok, err := VerifyWebhookSignature(sharedKey, signature, body)
	if err != nil {
		// handle error
		panic(err)
	}

	if !ok {
		// handle invalid signature
	}

```
## TODOs

- [ ] Add some mock tests, as coinbase-commerce does not currently have a sandbox
- [ ] Handle pagination properly
- [x] Add Webhook signature verification

License
----

MIT
