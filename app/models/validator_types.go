package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

// InputValidation - an interface for all "input submission" structs used for
// deserialization.  We pass in the request so that we can potentially get the
// context by the request from our context manager
type InputValidation interface {
	Validate(r *http.Request) error
}

var (
	// ErrInvalidName - error when we have an invalid name
	TypeErr     = errors.New("order type required")
	PriceErr    = errors.New("order price required")
)

func (t BookOrder) Validate(r *http.Request) error {
	// validate the name is not empty or missing
	if govalidator.IsNull(t.Side) {
		return TypeErr
	}
	if t.Price <= 0{
		return PriceErr
	}
	return nil
}
