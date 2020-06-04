package y

import (
	"errors"
	"fmt"
)

var (
	_ error                       = &Core{}
	_ interface{ Unwrap() error } = &Core{}
	_ interface{ Is(error) bool } = &Core{}
)

// Core is a general purpose error with a given identifier to help
// categorize and compare between different errors.
type Core struct {
	identity Identity
	err      error
}

// Fill sets the value of the core's error with a wrapped error.
func (c *Core) Fill(identity Identity, format string, args ...interface{}) {
	c.identity = identity
	c.err = fmt.Errorf(format, args...)
}

// Is compares other Core errors and checks their identifiers.
func (c *Core) Is(target error) bool {
	for {
		// Check if the error is an identity.
		var ident Identity
		if errors.As(target, &ident) {
			return ident == c.identity
		}
		// Check if the error is a core.
		tmpCore := &Core{}
		if errors.As(target, &tmpCore) {
			return tmpCore.identity == c.identity
		}
		// Try to unwrap the error.
		if target = errors.Unwrap(target); target == nil {
			return false
		}
	}
}

func (c *Core) Error() string {
	if c.err != nil {
		return c.err.Error()
	}

	return c.identity.Error()
}

// Identity returns the identity of the core.
func (c *Core) Identity() Identity {
	return c.identity
}

// Unwrap gives the internal error.
func (c *Core) Unwrap() error {
	return errors.Unwrap(c.err)
}

// Identity is a barebones string error, and can be paired with a core.
type Identity string

func (i Identity) Error() string {
	return string(i)
}

// NewError returns the identity of a new error.
func NewError(identity string) Identity {
	return Identity(identity)
}
