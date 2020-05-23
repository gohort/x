package x

import (
	"github.com/gohort/x/internal/y"
)

// NewError returns a new core error with an identifier.
// The core can be used as a standard error like when using errors.New or
// they can be used with the custom Errorf to categorize other errors.
func NewError(identity string) *y.Core {
	return y.NewError(identity)
}

// Errorf returns a new custom error with a particular type.
func Errorf(c *y.Core, format string, args ...interface{}) error {
	cc := *c
	cc.Fill(format, args...)
	return &cc
}
