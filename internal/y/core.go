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
	identity string
	err      error
}

// Fill sets the value of the core's error with a wrapped error.
func (c *Core) Fill(format string, args ...interface{}) {
	c.err = fmt.Errorf(format, args...)
}

// Is compares other Core errors and checks their identifiers.
func (c *Core) Is(target error) bool {
	if tCore, ok := target.(*Core); ok {
		if c.identity == tCore.identity {
			return true
		}
	}

	return false
}

func (c *Core) Error() string {
	if c.err != nil {
		return c.err.Error()
	}

	return c.identity
}

// Unwrap gives the internal error.
func (c *Core) Unwrap() error {
	return errors.Unwrap(c.err)
}

// NewError returns a core error with an identifier.
func NewError(identity string) *Core {
	return &Core{identity: identity}
}
