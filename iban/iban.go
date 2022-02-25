package iban

import (
	"context"

	"github.com/trtstm/iban-service/domain"
)

type IBANService struct {
}

// NewIBANService creates a new IBAN service.
func NewIBANService() *IBANService {
	return &IBANService{}
}

// ValidateIBAN validates the IBAN string and returns true if valid and false + an error if not.
func (s *IBANService) ValidateIBAN(_ context.Context, iban string) (bool, error) {
	_, err := domain.Parse(iban)

	if err != nil {
		return false, err
	}

	return true, nil
}
