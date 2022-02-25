package httpserver

import "context"

type ibanService interface {
	ValidateIBAN(context.Context, string) (bool, error)
}
