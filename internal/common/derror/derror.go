// Package derror defines internal error values to categorize the different
// types error semantics.
package derror

import (
	"errors"
	"fmt"
)

var (
	// NotFound indicates that a requested entity was not found (HTTP 404).
	NotFound = errors.New("not found")

	// InvalidArgument indicates that the input into the request is invalid in
	// some way (HTTP 400).
	InvalidArgument = errors.New("invalid argument")
)

type ErrResponse struct {
	Message string `json:"message"`
}

func Wrap(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}
