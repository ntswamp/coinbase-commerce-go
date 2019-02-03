package coinbase

import "time"

type WebhookEvent struct {
	ID           int       `json:"id"`
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
