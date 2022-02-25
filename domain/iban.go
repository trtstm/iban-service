package domain

import (
	"errors"
	"regexp"
	"strings"
)

type IBAN struct {
	countryCode string
	checkDigits uint
	bban        string
}

var countryCodeRe = regexp.MustCompile("^[a-zA-Z]{2}$")
var checkDigitsRe = regexp.MustCompile("^[0-9]{2}$")

// Parse parses a IBAN string
func Parse(iban string) (IBAN, error) {
	// Remove spaces.
	iban = strings.ReplaceAll(iban, " ", "")

	// Need atleast country code and check digits.
	if len(iban) < 5 || len(iban) > 34 {
		return IBAN{}, errors.New("invalid IBAN length")
	}

	countryCode := iban[0:2]
	if !countryCodeRe.MatchString(countryCode) {
		return IBAN{}, errors.New("invalid country code")
	}

	checkDigitsStr := iban[2:4]
	if !checkDigitsRe.MatchString(checkDigitsStr) {
		return IBAN{}, errors.New("invalid check digits")
	}

	return IBAN{}, nil
}
