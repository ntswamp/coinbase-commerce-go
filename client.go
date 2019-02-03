package coinbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const defaultCoinbaseURL = "https://api.commerce.coinbase.com"
const defaultAPIVersion = "2018-03-22"

type CoinbaseClient struct {
	apiKey     string
	apiVersion string
	apiBaseURL string

	httpClient *http.Client
}

func NewHttpClient(apiKey string) *CoinbaseClient {
	c := &http.Client{}
	c.Timeout = 3 * time.Second

	return &CoinbaseClient{
		apiKey:     apiKey,
		apiVersion: defaultAPIVersion,
		apiBaseURL: defaultCoinbaseURL,
		httpClient: c,
	}
}

func (c *CoinbaseClient) WithClientTimeout(numberOfMillis int) *CoinbaseClient {
	c.httpClient.Timeout = time.Duration(numberOfMillis) * time.Millisecond
	return c
}

func (c *CoinbaseClient) setDefaultHeaders(req *http.Request) {
	req.Header.Set("X-CC-Api-Key", c.apiKey)
	req.Header.Set("X-CC-Version", c.apiVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (c *CoinbaseClient) request(method, path string, body, result interface{}) (res *http.Response, err error) {
	parsedURL, err := url.Parse(c.apiBaseURL + path)
	if err != nil {
		return res, err
	}

	var data []byte
	rawBody := bytes.NewReader(make([]byte, 0))

	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return res, err
		}

		rawBody = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, parsedURL.String(), rawBody)

	if err != nil {
		return res, err
	}

	c.setDefaultHeaders(req)

	res, err = c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	// bd, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(bd), res.StatusCode)
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		responseError := ResponseError{
			HttpStatusCode: res.StatusCode,
		}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&responseError); err != nil {
			// TODO: If we could not decode the response error, we should probably still return a specific error, since it's an API error, with a malformed body.
			return res, err
		}

		return res, responseError
	}

	if result != nil {
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&result); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (c *CoinbaseClient) ListCheckouts() (checkouts []Checkout, err error) {
	var retrievedCheckouts ListCheckoutsResponse

	_, err = c.request("GET", "/checkouts", nil, &retrievedCheckouts)

	if err != nil {
		return checkouts, err
	}

	return retrievedCheckouts.Data, nil
}

func (c *CoinbaseClient) RetrieveCheckout(checkoutID string) (checkout Checkout, err error) {
	var retrievedCheckout RetrieveCheckoutResponse

	_, err = c.request("GET", "/checkouts/"+checkoutID, nil, &retrievedCheckout)

	if err != nil {
		return checkout, err
	}

	return retrievedCheckout.Data, nil
}

func (c *CoinbaseClient) CreateCheckout(ch Checkout) (checkout Checkout, err error) {
	var createdCheckout CreateCheckoutResponse
	_, err = c.request("POST", "/checkouts", ch, &createdCheckout)

	if err != nil {
		return checkout, err
	}

	return createdCheckout.Data, nil
}

func (c *CoinbaseClient) UpdateCheckout(ch Checkout) (checkout Checkout, err error) {
	var updatedCheckout UpdateCheckoutResponse
	_, err = c.request("PUT", "/checkouts/"+ch.ID, ch, &updatedCheckout)

	if err != nil {
		return checkout, err
	}

	return updatedCheckout.Data, nil
}

func (c *CoinbaseClient) DeleteCheckout(checkoutID string) error {
	_, err := c.request("DELETE", "/checkouts/"+checkoutID, nil, nil)

	return err
}

func (c *CoinbaseClient) ListCharges() (charges []Charge, err error) {
	var retrievedCharges ListChargesResponse

	_, err = c.request("GET", "/charges", nil, &retrievedCharges)

	if err != nil {
		return charges, err
	}

	return retrievedCharges.Data, nil
}

func (c *CoinbaseClient) RetrieveCharge(chargeID string) (charge Charge, err error) {
	var retrievedCharge RetrieveChargeResponse
	_, err = c.request("GET", "/charges/"+chargeID, nil, &retrievedCharge)

	if err != nil {
		return charge, err
	}

	return retrievedCharge.Data, nil

}

func (c *CoinbaseClient) CreateCharge(chargeRequest ChargeRequest) (charge Charge, err error) {
	var createdCharge CreateChargeResponse
	_, err = c.request("POST", "/charges", chargeRequest, &createdCharge)

	if err != nil {
		return charge, err
	}

	return createdCharge.Data, nil

}

func (c *CoinbaseClient) CancelCharge(chargeID string) (charge Charge, err error) {
	var retrievedCharge RetrieveChargeResponse
	targetURL := fmt.Sprintf("/charges/%s/cancel", chargeID)
	_, err = c.request("GET", targetURL, nil, &retrievedCharge)

	if err != nil {
		return charge, err
	}

	return retrievedCharge.Data, nil

}

func (c *CoinbaseClient) ResolveCharge(chargeID string) (charge Charge, err error) {
	var retrievedCharge RetrieveChargeResponse
	targetURL := fmt.Sprintf("/charges/%s/resolve", chargeID)
	_, err = c.request("GET", targetURL, nil, &retrievedCharge)

	if err != nil {
		return charge, err
	}

	return retrievedCharge.Data, nil

}
