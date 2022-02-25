package iban_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/trtstm/iban-service/iban"
)

func Test_iban_ValidateIBAN(t *testing.T) {
	service := iban.NewIBANService()

	for _, tc := range []struct {
		Name  string
		IBAN  string
		Valid bool
		Error error
	}{
		{
			Name:  "Too short IBAN",
			IBAN:  "BE12",
			Valid: false,
			Error: errors.New("invalid IBAN length"),
		},
		{
			Name:  "Too long IBAN",
			IBAN:  "BE777777777777777777777777777777777",
			Valid: false,
			Error: errors.New("invalid IBAN length"),
		},
		{
			Name:  "Wrong country code 1",
			IBAN:  "7777777777777777777777777777777",
			Valid: false,
			Error: errors.New("invalid country code"),
		},
		{
			Name:  "Wrong country code 2",
			IBAN:  "B7777777777777777777777777777777",
			Valid: false,
			Error: errors.New("invalid country code"),
		},
		{
			Name:  "Wrong check digits 1",
			IBAN:  "BEE777777777777777777777777777777",
			Valid: false,
			Error: errors.New("invalid check digits"),
		},
		{
			Name:  "Wrong check digits 2",
			IBAN:  "BE7E77777777777777777777777777777",
			Valid: false,
			Error: errors.New("invalid check digits"),
		},
		{
			Name:  "Wrong bban 1",
			IBAN:  "BE77 世界1231",
			Valid: false,
			Error: errors.New("invalid bban"),
		},
		{
			Name:  "Wrong bban 2",
			IBAN:  "BE77",
			Valid: false,
			Error: errors.New("invalid bban"),
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			valid, err := service.ValidateIBAN(tc.IBAN)
			fmt.Println(valid, err)
			if valid != tc.Valid || err.Error() != tc.Error.Error() {
				t.Errorf("expected %v, %v got %v, %v", tc.Valid, tc.Error, valid, err)
			}
		})
	}
}
