package x

import (
	"github.com/gohort/x/internal/y"
)

// NewError returns a new core error with an identifier.
// The core can be used as a standard error like when using errors.New or
// they can be used with the custom Errorf to categorize other errors.
func NewError(identity string) y.Identity {
	return y.NewError(identity)
}

// Errorf returns a new custom error with a particular type.
func Errorf(identity y.Identity, format string, args ...interface{}) error {
	cc := y.Core{}
	cc.Fill(identity, format, args...)
	return &cc
}

// Core is an export of the error core.
type Core = y.Core
