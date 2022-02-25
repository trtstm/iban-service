package domain

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type IBAN struct {
	countryCode string
	checkDigits uint
	bban        string
}

var countryCodeRe = regexp.MustCompile("^[a-zA-Z]{2}$")
var checkDigitsRe = regexp.MustCompile("^[0-9]{2}$")
var bbanRe = regexp.MustCompile("^[a-zA-Z0-9]{1,30}$")

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
	checkDigits, err := strconv.ParseUint(checkDigitsStr, 10, 32)
	if err != nil {
		return IBAN{}, errors.Wrap(err, "could not parse check digits as uint")
	}

	bban := iban[4:]
	if !bbanRe.MatchString(bban) {
		return IBAN{}, errors.New("invalid bban")
	}

	return IBAN{countryCode: countryCode, checkDigits: uint(checkDigits), bban: bban}, nil
}
