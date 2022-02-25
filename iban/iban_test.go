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
			Name:  "Valid IBAN 1",
			IBAN:  "NL89ABNA3729480081",
			Valid: true,
			Error: nil,
		},
		{
			Name:  "Valid IBAN 2",
			IBAN:  "AE240263856677841186492",
			Valid: true,
			Error: nil,
		},
		{
			Name:  "Valid IBAN 3",
			IBAN:  "SE9561918249885974448284",
			Valid: true,
			Error: nil,
		},
		{
			Name:  "Valid IBAN 4",
			IBAN:  "CH0989144485529931127",
			Valid: true,
			Error: nil,
		},
		{
			Name:  "Valid IBAN 5",
			IBAN:  "PS21AGRZ549376376519352845978",
			Valid: true,
			Error: nil,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			valid, err := service.ValidateIBAN(tc.IBAN)
			fmt.Println(valid, err)
			errStr := ""
			if err != nil {
				errStr = err.Error()
			}
			expectedErrorStr := ""
			if tc.Error != nil {
				expectedErrorStr = tc.Error.Error()
			}
			if valid != tc.Valid || errStr != expectedErrorStr {
				t.Errorf("expected %v, %v got %v, %v", tc.Valid, tc.Error, valid, err)
			}
		})
	}
}
