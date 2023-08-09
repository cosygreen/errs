// Package errs provides some special error types to be used with the log package.
//
// Features:
//   - error with stack trace
package errs

import (
	"errors"

	"github.com/tehsphinx/cstack"
)

// New creates a new error containing a stack.
func New(msg string) error {
	const skipStack = 3

	return stackError{
		//nolint:goerr113
		err:   errors.New(msg),
		stack: cstack.CallStack(skipStack),
	}
}
