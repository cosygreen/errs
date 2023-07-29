package errs

import (
	"errors"

	"github.com/tehsphinx/cstack"
)

// New creates a new error containing a stack.
func New(msg string) error {
	const skipStack = 3

	return stackErr{
		err:   errors.New(msg),
		stack: cstack.CallStack(skipStack),
	}
}
