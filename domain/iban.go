package domain

import (
	"math/big"
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

var countryCodeRe = regexp.MustCompile("^[A-Z]{2}$")
var checkDigitsRe = regexp.MustCompile("^[0-9]{2}$")
var bbanRe = regexp.MustCompile("^[A-Z0-9]{1,30}$")

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

	mod, err := calculateMOD97(countryCode + checkDigitsStr + bban)
	if err != nil {
		return IBAN{}, errors.Wrap(err, "could not calculate mod97")
	}

	if mod != 1 {
		return IBAN{}, errors.New("invalid check digits")
	}

	return IBAN{countryCode: countryCode, checkDigits: uint(checkDigits), bban: bban}, nil
}

func letterToDigits(r rune) string {
	// A in ascii is 65 and we want to map it to 10.
	return strconv.Itoa(int(r) - 55)
}

func calculateMOD97(s string) (uint, error) {
	// Move first 4 char to end.
	s = s[4:] + s[0:4]

	var result string

	// Convert to numbers.
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			result += string(r)
			continue
		}

		result += letterToDigits(r)
	}

	i := new(big.Int)

	_, ok := i.SetString(result, 10)
	if !ok {
		return 0, errors.New("could not set big int")
	}

	mod := new(big.Int)
	mod.Mod(i, big.NewInt(97))
	return uint(mod.Uint64()), nil
}
