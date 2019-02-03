package coinbase

// Error types returned by Coinbase Commerce API
const (
	NotFound            = "not_found"
	ParamRequired       = "param_required"
	ValidationError     = "validation_error"
	InvalidRequest      = "invalid_request"
	AuthenticationError = "authentication_error"
	RateLimitExceeded   = "rate_limit_exceeded"
	InternalServerError = "internal_server_error"
)

type ResponseError struct {
	HttpStatusCode int
	ReturnedError  struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

func (e ResponseError) Error() string {
	return e.ReturnedError.Message
}

func (e ResponseError) IsInvalidRequest() bool {
	return e.ReturnedError.Type == InvalidRequest
}

func (e ResponseError) IsValidationError() bool {
	return e.ReturnedError.Type == ValidationError
}

func (e ResponseError) IsAuthenticationError() bool {
	return e.ReturnedError.Type == AuthenticationError
}

func (e ResponseError) IsRateLimitExceeded() bool {
	return e.ReturnedError.Type == RateLimitExceeded
}

func (e ResponseError) IsInternalServerError() bool {
	// Probably status code is enough?
	return e.HttpStatusCode == 500 || e.ReturnedError.Type == InternalServerError
}

func (e ResponseError) IsNotFound() bool {
	return e.HttpStatusCode == 404 || e.ReturnedError.Type == NotFound
}
