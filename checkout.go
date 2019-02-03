package coinbase

// Pricing Type Options
const (
	FixedPrice = "fixed_price"
	NoPrice    = "no_price"
)

type LocalPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Checkout struct {
	ID            string     `json:"id,omitempty"`
	Resource      string     `json:"resource,omitempty"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	LogoURL       string     `json:"logo_url,omitempty"`
	RequestedInfo []string   `json:"requested_info,omitempty"` // like "name" and "email"
	PricingType   string     `json:"pricing_type"`
	LocalPrice    LocalPrice `json:"local_price"`
}

type ListCheckoutsParams struct {
	Pagination *Pagination
}

type ListCheckoutsResponse struct {
	Pagination *Pagination `json:"pagination"`
	Data       []Checkout  `json:"data"`
}

type RetrieveCheckoutResponse struct {
	Data Checkout `json:"data"`
}

type CreateCheckoutResponse struct {
	Data Checkout `json:"data"`
}

type UpdateCheckoutResponse struct {
	Data Checkout `json:"data"`
}
