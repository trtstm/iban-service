package http

type validateIBANResponse struct {
	// Whether IBAN is valid or not.
	Valid bool `json:"valid"`

	// If not valid, this is the reason why.
	Reason string `json:"reason,omitempty"`
}
