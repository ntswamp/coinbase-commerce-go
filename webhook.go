package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type WebhookEvent struct {
	ID           string       `json:"id"`
	ScheduledFor time.Time `json:"scheduled_for"`
	Event        struct {
		ID         string    `json:"id"`
		Resource   string    `json:"resource"`
		Type       string    `json:"type"`
		APIVersion string    `json:"api_version"`
		CreatedAt  time.Time `json:"created_at"`
		Data       struct {
			Code        string    `json:"code"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			HostedURL   string    `json:"hosted_url"`
			CreatedAt   time.Time `json:"created_at"`
			ExpiresAt   time.Time `json:"expires_at"`
			Timeline    []struct {
				Time   time.Time `json:"time"`
				Status string    `json:"status"`
			} `json:"timeline"`
			Metadata struct {
			} `json:"metadata"`
			PricingType string        `json:"pricing_type"`
			Payments    []interface{} `json:"payments"`
			Addresses   struct {
				Bitcoin  string `json:"bitcoin"`
				Ethereum string `json:"ethereum"`
			} `json:"addresses"`
		} `json:"data"`
	} `json:"event"`
}

func verify(sharedKey, signature, payload []byte) bool {
	mac := hmac.New(sha256.New, sharedKey)
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(expectedMAC, signature)
}

func VerifyWebhookSignatureFromRequest(sharedKey string, r *http.Request) (bool, error) {
	signature := r.Header.Get("X-CC-Webhook-Signature")

	if signature == "" {
		return false, errors.New("webhook signature is missing from the request")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, err
	}
	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	return VerifyWebhookSignature(sharedKey, signature, body)
}

func VerifyWebhookSignature(sharedKey, signature string, body []byte) (bool, error) {
	return verify([]byte(sharedKey), []byte(signature), body), nil
}
