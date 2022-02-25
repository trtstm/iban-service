package domain

import (
	"errors"
	"fmt"
	"strings"
)

type bban string
type countryCode string

type IBAN struct {
	countryCode countryCode
	checkDigits uint
	bban        bban
}

// Parse parses a IBAN string
func Parse(iban string) (IBAN, error) {
	// Remove spaces.
	iban = strings.ReplaceAll(iban, " ", "")

	// Need atleast country code and check digits.
	if len(iban) < 5 {
		return IBAN{}, errors.New("invalid IBAN length")
	}

	countryCode := countryCode(iban[0:2])
	fmt.Println(countryCode)

	return IBAN{}, nil
}
